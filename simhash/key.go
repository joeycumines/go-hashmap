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

// Key is an interface based on the core Java Object, which forms the basis for the hashmap implementation.
// See https://www.sitepoint.com/how-to-implement-javas-hashcode-correctly/ or look around for details on implementing
// this.
type Key interface {
	Hash() int
	Equals(other interface{}) bool
}
