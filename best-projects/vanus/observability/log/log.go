// Copyright 2022 Linkall Inc.
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

package log

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

var lg zerolog.Logger

func init() {
	level := os.Getenv("VANUS_LOG_LEVEL")
	var lvl zerolog.Level
	switch strings.ToLower(level) {
	case "debug":
		lvl = zerolog.DebugLevel
	case "info":
		lvl = zerolog.InfoLevel
	case "error":
		lvl = zerolog.ErrorLevel
	case "fatal":
		lvl = zerolog.FatalLevel
	default:
		lvl = zerolog.WarnLevel
	}
	lg = zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Caller().Logger().Level(lvl)
}

func SetOutput(w io.Writer) {
	lg = lg.Output(w)
}

func With() zerolog.Context {
	return lg.With()
}

func Debug(_ ...context.Context) *zerolog.Event {
	return lg.Debug()
}

func Info(_ ...context.Context) *zerolog.Event {
	return lg.Info()
}

func Warn(_ ...context.Context) *zerolog.Event {
	return lg.Warn()
}

func Error(_ ...context.Context) *zerolog.Event {
	return lg.Error()
}
