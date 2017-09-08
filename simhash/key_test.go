/*
   Copyright 2017 Joseph Cumines

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

package simhash

import "strconv"

type testKeyInt int

func (k testKeyInt) Hash() int {
	return int(k)
}

func (k testKeyInt) Equals(other interface{}) bool {
	return int(k) == int(other.(testKeyInt))
}

type testKeyStruct struct {
	hash int
	val  int
}

func (k testKeyStruct) Hash() int {
	return k.hash
}

func (k testKeyStruct) Equals(other interface{}) bool {
	return k == other.(testKeyStruct)
}

func (k testKeyStruct) String() string {
	return strconv.Itoa(k.val)
}
