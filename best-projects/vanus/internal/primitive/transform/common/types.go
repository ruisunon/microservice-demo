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

package common

type Type uint8

const (
	String Type = iota
	Float
	Int
	Bool
	Object
	Array
	StringArray
	Any
)

func (t Type) String() string {
	switch t {
	case String:
		return "string"
	case Float:
		return "float"
	case Int:
		return "int"
	case Bool:
		return "bool"
	case Object:
		return "map"
	case Array:
		return "array"
	case Any:
		return "any"
	case StringArray:
		return "stringArray"
	}
	return "unknown"
}

func TypePtr(t Type) *Type {
	return &t
}

func (t Type) IsSameType(val interface{}) bool {
	return TypeFromVal(val) == t
}

func TypeFromVal(val interface{}) Type {
	switch val.(type) {
	case string:
		return String
	case float64:
		return Float
	case int:
		return Int
	case bool:
		return Bool
	case map[string]interface{}:
		return Object
	case []string:
		return StringArray
	case []interface{}:
		return Array
	}
	return Any
}
