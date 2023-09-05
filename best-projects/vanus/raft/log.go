// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package raft

import (
	"fmt"
	"log"

	pb "github.com/vanus-labs/vanus/raft/raftpb"
)

type raftLog struct {
	// storage contains all stable entries since the last snapshot.
	storage Storage
	// keeper writes entries to the stable storage.
	keeper Keeper

	inflight inflight

	// unstable contains all unstable entries and snapshot.
	// they will be saved into storage.
	unstable unstable

	// persisting is the next log position that will be persisted to storage.
	// Invariant: unstable.offset <= persisting
	persisting uint64
	// committed is the highest log position that is known to be in
	// stable storage on a quorum of nodes.
	// Invariant: committed < unstable.offset + len(unstable.entries)
	committed uint64
	// Invariant: localCommitted = min(committed, unstable.offset)
	localCommitted uint64
	// applying is the highest log position that the application has
	// been instructed to be applying to its state machine.
	// Invariant: applying <= localCommitted
	applying uint64
	// applied is the highest log position that the application has
	// been instructed to apply to its state machine.
	// Invariant: applied <= applying
	applied uint64
	// compacted is the highest log position that the application can
	// delete safety.
	// Invariant: compacted <= applied
	compacted uint64

	logger Logger

	// maxNextEntsSize is the maximum number aggregate byte size of the messages
	// returned from calls to nextEnts.
	maxNextEntsSize uint64
}

// newLog returns log using the given storage and default options. It
// recovers the log to the state that it just commits and applies the
// latest snapshot.
func newLog(storage Storage, keeper Keeper, logger Logger) *raftLog {
	return newLogWithSize(storage, keeper, logger, noLimit)
}

// newLogWithSize returns a log using the given storage and max
// message size.
func newLogWithSize(storage Storage, keeper Keeper, logger Logger, maxNextEntsSize uint64) *raftLog {
	if storage == nil {
		log.Panic("storage must not be nil")
	}
	log := &raftLog{
		storage:         storage,
		keeper:          keeper,
		logger:          logger,
		maxNextEntsSize: maxNextEntsSize,
	}
	firstIndex, err := storage.FirstIndex()
	if err != nil {
		panic(err) // TODO(bdarnell)
	}
	lastIndex, err := storage.LastIndex()
	if err != nil {
		panic(err) // TODO(bdarnell)
	}
	log.unstable.offset = lastIndex + 1
	log.unstable.logger = logger
	log.persisting = lastIndex + 1
	// Initialize our committed and applied pointers to the time of the last compaction.
	log.committed = firstIndex - 1
	log.localCommitted = firstIndex - 1
	log.applying = firstIndex - 1
	log.applied = firstIndex - 1
	log.compacted = firstIndex - 1

	return log
}

func (l *raftLog) String() string {
	return fmt.Sprintf("committed=%d, applied=%d, unstable.offset=%d, len(unstable.Entries)=%d",
		l.committed, l.applied, l.unstable.offset, len(l.unstable.entries))
}

// maybeAppend returns false if the entries cannot be appended. Otherwise, it returns true.
func (l *raftLog) maybeAppend(index, logTerm, committed uint64, ents ...pb.Entry) (ok bool) {
	if !l.matchTerm(index, logTerm) {
		return false
	}

	li := index + uint64(len(ents))
	ci := l.findConflict(ents)
	switch {
	case ci == 0:
	case ci <= l.committed:
		l.logger.Panicf("entry %d conflict with committed entry [committed(%d)]", ci, l.committed)
	default:
		offset := index + 1
		if ci-offset > uint64(len(ents)) {
			l.logger.Panicf("index, %d, is out of range [%d]", ci-offset, len(ents))
		}
		l.append(ents[ci-offset:]...)
	}
	l.commitTo(min(committed, li))
	return true
}

func (l *raftLog) append(ents ...pb.Entry) uint64 {
	if len(ents) == 0 {
		return l.lastIndex()
	}
	if prev := ents[0].Index - 1; prev < l.committed {
		l.logger.Panicf("prev(%d) is out of range [committed(%d)]", prev, l.committed)
	}

	li, truncated := l.unstable.truncateAndAppend(ents)
	start := ents[0].Index
	if truncated {
		l.inflight.truncateFrom(start)
	}
	// Reset pending when any entry being persisted is truncated.
	if l.persisting > start {
		l.persisting = start
	}

	// TODO(james.yin): limiting
	l.keeper.TruncateAndAppend(ents)
	l.persisting = li + 1

	return li
}

// findConflict finds the index of the conflict.
// It returns the first pair of conflicting entries between the existing
// entries and the given entries, if there are any.
// If there is no conflicting entries, and the existing entries contains
// all the given entries, zero will be returned.
// If there is no conflicting entries, but the given entries contains new
// entries, the index of the first new entry will be returned.
// An entry is considered to be conflicting if it has the same index but
// a different term.
// The index of the given entries MUST be continuously increasing.
func (l *raftLog) findConflict(ents []pb.Entry) uint64 {
	for i := range ents {
		ne := &ents[i]
		if !l.matchTerm(ne.Index, ne.Term) {
			if ne.Index <= l.lastIndex() {
				l.logger.Infof("found conflict at index %d [existing term: %d, conflicting term: %d]",
					ne.Index, l.zeroTermOnErrCompacted(l.term(ne.Index)), ne.Term)
			}
			return ne.Index
		}
	}
	return 0
}

// findConflictByTerm takes an (index, term) pair (indicating a conflicting log
// entry on a leader/follower during an append) and finds the largest index in
// log l with a term <= `term` and an index <= `index`. If no such index exists
// in the log, the log's first index is returned.
//
// The index provided MUST be equal to or less than l.lastIndex(). Invalid
// inputs log a warning and the input index is returned.
func (l *raftLog) findConflictByTerm(index uint64, term uint64) uint64 {
	if li := l.lastIndex(); index > li {
		// NB: such calls should not exist, but since there is a straightfoward
		// way to recover, do it.
		//
		// It is tempting to also check something about the first index, but
		// there is odd behavior with peers that have no log, in which case
		// lastIndex will return zero and firstIndex will return one, which
		// leads to calls with an index of zero into this method.
		l.logger.Warningf("index(%d) is out of range [0, lastIndex(%d)] in findConflictByTerm",
			index, li)
		return index
	}
	for {
		logTerm, err := l.term(index)
		if logTerm <= term || err != nil {
			break
		}
		index--
	}
	return index
}

// hasPendingSnapshot returns if there is pending snapshot waiting for applying.
func (l *raftLog) hasPendingSnapshot() bool {
	return l.unstable.snapshot != nil && !IsEmptySnap(*l.unstable.snapshot)
}

func (l *raftLog) snapshot() (pb.Snapshot, error) {
	if l.unstable.snapshot != nil {
		return *l.unstable.snapshot, nil
	}
	return l.storage.Snapshot()
}

func (l *raftLog) firstIndex() uint64 {
	if i, ok := l.unstable.maybeFirstIndex(); ok {
		return i
	}
	index, err := l.storage.FirstIndex()
	if err != nil {
		panic(err) // TODO(bdarnell)
	}
	return index
}

func (l *raftLog) lastIndex() uint64 {
	if i, ok := l.unstable.maybeLastIndex(); ok {
		return i
	}
	return l.unstable.offset - 1
}

func (l *raftLog) stableLastIndex() uint64 { //nolint:unused // ok
	i, err := l.storage.LastIndex()
	if err != nil {
		panic(err) // TODO(james.yin)
	}
	return i
}

func (l *raftLog) compactTo(tocompact uint64) {
	if l.applied < tocompact {
		l.logger.Debugf("tocompact(%d) is less then applied(%d).", tocompact, l.applied)
		tocompact = l.applied
	}
	if l.compacted < tocompact {
		l.compacted = tocompact
		l.keeper.CompactTo(tocompact)
	}
}

func (l *raftLog) appliedTo(i uint64) {
	if i == 0 {
		return
	}
	if l.applying < i || i < l.applied {
		l.logger.Panicf("applied(%d) is out of range [prevApplied(%d), applying(%d)]", i, l.applied, l.applying)
	}
	l.applied = i
}

func (l *raftLog) applyingTo(i uint64) {
	if i == 0 {
		return
	}
	if l.localCommitted < i || i < l.applying {
		l.logger.Panicf("applied(%d) is out of range [prevApplying(%d), localCommitted(%d)]", i, l.applying, l.localCommitted)
	}
	l.applying = i
}

func (l *raftLog) localCommitTo(tocommit uint64) {
	if tocommit >= l.unstable.offset {
		tocommit = l.unstable.offset - 1
	}
	// never decrease commit
	if l.localCommitted < tocommit {
		// if li := l.stableLastIndex(); li < tocommit {
		// 	l.logger.Panicf("tocommit(%d) is out of range [lastIndex(%d)]. Was the raft log corrupted, truncated, or lost?", tocommit, li)
		// }
		l.localCommitted = tocommit
		l.inflight.commitTo(tocommit)
		l.keeper.CommitTo(tocommit)

		// TODO(james.yin): limiting
		ents, err := l.slice(l.applying+1, tocommit+1, noLimit)
		if err != nil {
			l.logger.Panicf("unexpected error when getting unapplied entries (%v)", err)
		}
		l.keeper.Apply(ents)
		l.applying = tocommit
	}
}

func (l *raftLog) commitTo(tocommit uint64) {
	// never decrease commit
	if l.committed < tocommit {
		if li := l.lastIndex(); li < tocommit {
			l.logger.Panicf("tocommit(%d) is out of range [lastIndex(%d)]. Was the raft log corrupted, truncated, or lost?", tocommit, li)
		}
		l.committed = tocommit
	}
	l.localCommitTo(tocommit)
}

func (l *raftLog) persistingTo(i, t uint64) {
	gt, ok := l.unstable.maybeTerm(i)
	if !ok {
		return
	}
	// if i < offset, term is matched with the snapshot
	// only update the pending if term is matched with an unstable entry.
	if gt == t && i >= l.unstable.offset {
		l.persisting = i + 1
	}
}

func (l *raftLog) stableTo(i, t uint64) bool {
	return l.unstable.stableTo(i, t)
}

func (l *raftLog) stableSnapTo(i uint64) {
	if l.unstable.stableSnapTo(i) {
		l.compactTo(i)
	}
}

func (l *raftLog) lastTerm() uint64 {
	t, err := l.term(l.lastIndex())
	if err != nil {
		l.logger.Panicf("unexpected error when getting the last term (%v)", err)
	}
	return t
}

func (l *raftLog) term(i uint64) (uint64, error) {
	// Check the unstable log first, even before computing the valid term range,
	// which may need to access stable Storage. If we find the entry's term in
	// the unstable log, we know it was in the valid range.
	if t, ok := l.unstable.maybeTerm(i); ok {
		return t, nil
	}

	// The valid term range is [index of dummy entry, last index].
	dummyIndex := l.firstIndex() - 1
	if i < dummyIndex || i > l.lastIndex() {
		// TODO: return an error instead?
		return 0, nil
	}

	t, err := l.storage.Term(i)
	if err == nil {
		return t, nil
	}
	if err == ErrCompacted || err == ErrUnavailable {
		return 0, err
	}
	panic(err) // TODO(bdarnell)
}

func (l *raftLog) stableTerm(i uint64) (uint64, error) {
	// the valid term range is [index of dummy entry, last index]
	dummyIndex := l.firstIndex() - 1
	if i < dummyIndex || i > l.stableLastIndex() {
		// TODO: return an error instead?
		return 0, nil
	}

	t, err := l.storage.Term(i)
	if err == nil {
		return t, nil
	}
	if err == ErrCompacted || err == ErrUnavailable {
		return 0, err
	}
	panic(err) // TODO(bdarnell)
}

func (l *raftLog) entries(i, maxsize uint64) ([]pb.Entry, error) {
	if i > l.lastIndex() {
		return nil, nil
	}
	return l.slice(i, l.lastIndex()+1, maxsize)
}

// allEntries returns all entries in the log.
func (l *raftLog) allEntries() []pb.Entry {
	ents, err := l.entries(l.firstIndex(), noLimit)
	if err == nil {
		return ents
	}
	if err == ErrCompacted { // try again if there was a racing compaction
		return l.allEntries()
	}
	// TODO (xiangli): handle error?
	panic(err)
}

// isUpToDate determines if the given (lastIndex,term) log is more up-to-date
// by comparing the index and term of the last entries in the existing logs.
// If the logs have last entries with different terms, then the log with the
// later term is more up-to-date. If the logs end with the same term, then
// whichever log has the larger lastIndex is more up-to-date. If the logs are
// the same, the given log is up-to-date.
func (l *raftLog) isUpToDate(lasti, term uint64) bool {
	return term > l.lastTerm() || (term == l.lastTerm() && lasti >= l.lastIndex())
}

func (l *raftLog) matchTerm(i, term uint64) bool {
	t, err := l.term(i)
	if err != nil {
		return false
	}
	return t == term
}

func (l *raftLog) maybeCommit(maxIndex, term uint64) bool {
	if maxIndex > l.committed /* && l.zeroTermOnErrCompacted(l.term(maxIndex)) == term */ {
		l.commitTo(maxIndex)
		return true
	}
	return false
}

func (l *raftLog) maybeCompact(i uint64) {
	compacted := min(l.applied, i)
	if compacted > l.compacted {
		l.compactTo(compacted)
	}
}

func (l *raftLog) restore(s pb.Snapshot) {
	l.logger.Infof("log [%s] starts to restore snapshot [index: %d, term: %d]", l, s.Metadata.Index, s.Metadata.Term)
	l.committed = s.Metadata.Index
	l.localCommitted = s.Metadata.Index
	// NOTE: applied and compacted will be reset in raft.advance().
	l.unstable.restore(s)
	l.persisting = l.unstable.offset
}

// slice returns a slice of log entries from lo through hi-1, inclusive.
func (l *raftLog) slice(lo, hi, maxSize uint64) ([]pb.Entry, error) {
	err := l.mustCheckOutOfBounds(lo, hi)
	if err != nil {
		return nil, err
	}
	if lo == hi {
		return nil, nil
	}
	var ents []pb.Entry
	if lo < l.unstable.offset {
		storedEnts, err := l.storage.Entries(lo, min(hi, l.unstable.offset), maxSize)
		switch {
		case err == ErrCompacted: //nolint:errorlint // it's ok
			return nil, err
		case err == ErrUnavailable: //nolint:errorlint // it's ok
			l.logger.Panicf("entries[%d:%d) is unavailable from storage", lo, min(hi, l.unstable.offset))
		case err != nil:
			panic(err) // TODO(bdarnell)
		}

		// check if ents has reached the size limitation
		if uint64(len(storedEnts)) < min(hi, l.unstable.offset)-lo {
			return storedEnts, nil
		}

		ents = storedEnts
	}
	if hi > l.unstable.offset {
		unstable := l.unstable.slice(max(lo, l.unstable.offset), hi)
		if len(ents) > 0 {
			combined := make([]pb.Entry, len(ents)+len(unstable))
			n := copy(combined, ents)
			copy(combined[n:], unstable)
			ents = combined
		} else {
			ents = unstable
		}
	}
	return limitSize(ents, maxSize), nil
}

// l.firstIndex <= lo <= hi <= l.firstIndex + len(l.entries)
func (l *raftLog) mustCheckOutOfBounds(lo, hi uint64) error {
	if lo > hi {
		l.logger.Panicf("invalid slice %d > %d", lo, hi)
	}

	// fi := l.firstIndex()
	// if lo < fi {
	// 	return ErrCompacted
	// }

	li := l.lastIndex()
	if hi > li+1 {
		l.logger.Panicf("slice[%d,%d) out of bound [unknown,%d]", lo, hi /*fi,*/, li)
	}

	return nil
}

func (l *raftLog) zeroTermOnErrCompacted(t uint64, err error) uint64 {
	if err == nil {
		return t
	}
	if err == ErrCompacted { //nolint:errorlint // it's ok
		return 0
	}
	l.logger.Panicf("unexpected error (%v)", err)
	return 0
}
