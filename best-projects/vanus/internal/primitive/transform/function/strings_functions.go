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

package function

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/vanus-labs/vanus/internal/primitive/transform/common"
)

var JoinFunction = function{
	name:         "JOIN",
	fixedArgs:    []common.Type{common.String, common.StringArray},
	variadicArgs: common.TypePtr(common.StringArray),
	fn: func(args []interface{}) (interface{}, error) {
		sep, _ := args[0].(string)
		var sb strings.Builder
		for i := 1; i < len(args)-1; i++ {
			sb.WriteString(strings.Join(args[i].([]string), sep))
			sb.WriteString(sep)
		}
		sb.WriteString(strings.Join(args[len(args)-1].([]string), sep))
		return sb.String(), nil
	},
}

var UpperFunction = function{
	name:      "UPPER_CASE",
	fixedArgs: []common.Type{common.String},
	fn: func(args []interface{}) (interface{}, error) {
		return strings.ToUpper(args[0].(string)), nil
	},
}

var LowerFunction = function{
	name:      "LOWER_CASE",
	fixedArgs: []common.Type{common.String},
	fn: func(args []interface{}) (interface{}, error) {
		return strings.ToLower(args[0].(string)), nil
	},
}

var AddPrefixFunction = function{
	name:      "ADD_PREFIX",
	fixedArgs: []common.Type{common.String, common.String},
	fn: func(args []interface{}) (interface{}, error) {
		return args[1].(string) + args[0].(string), nil
	},
}

var AddSuffixFunction = function{
	name:      "ADD_SUFFIX",
	fixedArgs: []common.Type{common.String, common.String},
	fn: func(args []interface{}) (interface{}, error) {
		return args[0].(string) + args[1].(string), nil
	},
}

var SplitWithSepFunction = function{
	name:         "SPLIT_WITH_SEP",
	fixedArgs:    []common.Type{common.String, common.String},
	variadicArgs: common.TypePtr(common.Int),
	fn: func(args []interface{}) (interface{}, error) {
		s, _ := args[0].(string)
		sep, _ := args[1].(string)
		if len(args) == 2 {
			return strings.Split(s, sep), nil
		}
		return strings.SplitN(s, sep, args[2].(int)), nil
	},
}

var ReplaceBetweenPositionsFunction = function{
	name:      "REPLACE_BETWEEN_POSITIONS",
	fixedArgs: []common.Type{common.String, common.Int, common.Int, common.String},
	fn: func(args []interface{}) (interface{}, error) {
		path, _ := args[0].(string)
		startPosition, _ := args[1].(int)
		endPosition, _ := args[2].(int)
		targetValue, _ := args[3].(string)
		if startPosition >= len(path) {
			return nil, fmt.Errorf("start position must be less than the length of the string")
		}
		if endPosition >= len(path) {
			return nil, fmt.Errorf("end position must be less than the length of the string")
		}
		if startPosition >= endPosition {
			return nil, fmt.Errorf("start position must be less than end position")
		}
		return path[:startPosition] + targetValue + path[endPosition:], nil
	},
}

var CapitalizeSentence = function{
	name:      "CAPITALIZE_SENTENCE",
	fixedArgs: []common.Type{common.String},
	fn: func(args []interface{}) (interface{}, error) {
		value, _ := args[0].(string)
		if len(value) == 0 {
			return value, nil
		}
		if len(value) == 1 {
			return strings.ToUpper(string(value[0])), nil
		}
		return strings.ToUpper(string(value[0])) + value[1:], nil
	},
}

var ReplaceBetweenDelimitersFunction = function{
	name:      "REPLACE_BETWEEN_DELIMITERS",
	fixedArgs: []common.Type{common.String, common.String, common.String, common.String},
	fn: func(args []interface{}) (interface{}, error) {
		value, _ := args[0].(string)
		startDelimiter, _ := args[1].(string)
		endDelimiter, _ := args[2].(string)
		replaceValue, _ := args[3].(string)
		if startDelimiter == "" || endDelimiter == "" {
			return nil, fmt.Errorf("start or end delemiter is empty")
		}
		startIndex := strings.Index(value, startDelimiter)
		if startIndex < 0 {
			return nil, fmt.Errorf("start delemiter is not exist")
		}
		index := startIndex + len(startDelimiter)
		endIndex := strings.Index(value[index:], endDelimiter)
		if endIndex < 0 {
			return nil, fmt.Errorf("end delemiter is not exist")
		}
		return value[:startIndex] + replaceValue + value[index+endIndex+len(endDelimiter):], nil

	},
}

var CapitalizeWord = function{
	name:      "CAPITALIZE_WORD",
	fixedArgs: []common.Type{common.String},
	fn: func(args []interface{}) (interface{}, error) {
		value, _ := args[0].(string)
		rs := []rune(value)
		inWord := false
		for i, r := range rs {
			if !unicode.IsSpace(r) {
				if !inWord {
					rs[i] = unicode.ToTitle(r)
				}
				inWord = true
			} else {
				inWord = false
			}
		}
		return string(rs), nil
	},
}

var SplitFromStart = function{
	name:      "SPLIT_FROM_START",
	fixedArgs: []common.Type{common.String, common.Int},
	fn: func(args []interface{}) (interface{}, error) {
		value, _ := args[0].(string)
		splitPosition, _ := args[1].(int)

		if splitPosition <= 0 {
			return nil, fmt.Errorf("split position must be more than zero")
		}

		if splitPosition >= len(value) {
			return []string{value, ""}, nil
		}

		result := []string{value[:splitPosition], value[splitPosition:]}
		return result, nil
	},
}
