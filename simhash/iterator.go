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

// Iterator provides an interface for iterating over the key-pairs of a map. It's value should initially start as nil,
// or uninitialized, and on the first call to either Next or Previous should determine the initial direction.
// After being initialized, Next (or Previous) can be used increment or de-increment the pointer, which will always
// stop at the last (or first) item in the iteration. The iterator can change direction, but will not move for the
// first call to the step in the opposite direction.
type Iterator interface {
	// Pair will return the current pair in the iteration, initially nil.
	Pair() Pair

	// Value will return the current value in the iteration, initially nil.
	Value() Value

	// Key will return the current key in the iteration, initially nil.
	Key() Key

	// Next moves the iterator in the forwards direction, or reverses the iterator (and doesn't move) if it is
	// currently iterating backwards, and will return false if there were no more items left to iterate.
	Next() bool

	// Previous moves the iterator in the backwards direction, or reverses the iterator (and doesn't move) if it is
	// currently iterating forwards, and will return false if there were no more items left to iterate.
	Previous() bool
}

type iterator struct {
	m        *hashMap
	hList    []int
	pair     Pair
	h        int
	i        int
	forwards bool
	active   bool
}

func (it *iterator) Pair() Pair {
	return it.pair
}

func (it *iterator) Value() Value {
	if nil == it.pair {
		return nil
	}
	return it.pair.Value()
}

func (it *iterator) Key() Key {
	if nil == it.pair {
		return nil
	}
	return it.pair.Key()
}

func (it *iterator) increment(forwards bool) bool {
	if 0 == len(it.hList) {
		return false
	}
	// Find if an index is outside the bounds of a slice.
	done := func(index, size int) bool {
		return 0 >= size || index >= size || index < 0
	}
	inc := 1
	if false == forwards {
		inc = -1
	}
	// Resets it.i to either the start or the end of the relevant box, depending on the direction of the increment.
	resetI := func() {
		it.i = 0
		if inc < 0 && false == done(it.h, len(it.hList)) {
			if hBox, ok := it.m.m[it.hList[it.h]]; true == ok && nil != hBox {
				it.i = len(hBox) - 1
			}
		}
	}
	resetH := func() {
		it.h = 0
		if inc < 0 {
			it.h = len(it.hList) - 1
		}
		resetI()
	}
	if false == it.active {
		it.active = true
		it.forwards = forwards
		resetH()
	}
	// if we changed directions, we need to increment until back in the correct range (if we are not in range).
	if forwards != it.forwards {
		if true == done(it.h, len(it.hList)) {
			resetH()
		} else {
			it.i += inc
		}
	}
	it.forwards = forwards
	for false == done(it.h, len(it.hList)) {
		hBox, ok := it.m.m[it.hList[it.h]]
		if false == ok || nil == hBox || true == done(it.i, len(hBox)) {
			it.h += inc
			resetI()
			continue
		}
		for false == done(it.i, len(hBox)) {
			pair := hBox[it.i]
			it.i += inc
			if nil == pair {
				continue
			}
			it.pair = pair
			return true
		}
	}
	return false
}

func (it *iterator) Next() bool {
	return it.increment(true)
}

func (it *iterator) Previous() bool {
	return it.increment(false)
}
