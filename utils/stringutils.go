/*
   Copyright 2020 The Compose Specification Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package utils

import (
	"strconv"
	"strings"
)

// StringContains check if an array contains a specific value
func StringContains(array []string, needle string) bool {
	for _, val := range array {
		if val == needle {
			return true
		}
	}
	return false
}

// StringToBool converts a string to a boolean ignoring errors
func StringToBool(s string) bool {
	b, _ := strconv.ParseBool(strings.ToLower(strings.TrimSpace(s)))
	return b
}
