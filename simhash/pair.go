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

// Pair represents a key-value relationship.
type Pair interface {
	Key() Key
	Value() Value
}

type pair struct {
	key   Key
	value Value
}

func (p pair) Key() Key {
	return p.key
}

func (p pair) Value() Value {
	return p.value
}

// NewPair creates a key-value pair.
func NewPair(key Key, value Value) Pair {
	return pair{key, value}
}
