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

import (
	"testing"
)

func TestIterator(t *testing.T) {
	m := genTestStructureHashMap()
	it := m.Iterator()

	pairList := make([]Pair, 0)
	valueList := make([]Value, 0)
	keyList := make([]Key, 0)

	comparePair := func(a, b Pair) bool {
		if a.Value().(int) != b.Value().(int) {
			return false
		}
		if nil == a.Key() && nil == b.Key() {
			return true
		}
		if nil == a.Key() || nil == b.Key() {
			return false
		}
		if a.Key().Hash() != a.Key().Hash() {
			return false
		}
		return a.Key().Equals(b.Key())
	}

	verifyLists := func() {
		if 10 != len(pairList) || 10 != len(valueList) || 10 != len(keyList) {
			t.Fatalf("%v", pairList)
		}
		// the actual => expected indexes, as found
		found := make(map[int]int)
		expected := m.Pairs()
		for i, pair := range expected {
			for ia, pa := range pairList {
				if _, exists := found[ia]; false == exists && true == comparePair(pair, pa) {
					if valueList[ia].(int) != pair.Value().(int) {
						t.Fatal()
					}
					if (nil == keyList[ia] && nil != pair.Key()) || (nil != keyList[ia] && nil == pair.Key()) {
						t.Fatal()
					}
					if nil != keyList[ia] && nil != pair.Key() && false == keyList[ia].Equals(pair.Key()) {
						t.Fatal()
					}
					found[ia] = i
					break;
				}
			}
		}
		if 10 != len(found) {
			t.Fatal()
		}
		pairList = make([]Pair, 0)
		valueList = make([]Value, 0)
		keyList = make([]Key, 0)
	}

	// test we can go forwards
	for true == it.Next() {
		pairList = append(pairList, it.Pair())
		valueList = append(valueList, it.Value())
		keyList = append(keyList, it.Key())
	}
	verifyLists()

	// test we can go backwards
	for true == it.Previous() {
		pairList = append(pairList, it.Pair())
		valueList = append(valueList, it.Value())
		keyList = append(keyList, it.Key())
	}
	verifyLists()

	// and forwards again
	for true == it.Next() {
		pairList = append(pairList, it.Pair())
		valueList = append(valueList, it.Value())
		keyList = append(keyList, it.Key())
	}
	verifyLists()

	// and backwards again
	for true == it.Previous() {
		pairList = append(pairList, it.Pair())
		valueList = append(valueList, it.Value())
		keyList = append(keyList, it.Key())
	}
	verifyLists()

	// test the inner values as we step manually
	itC := m.Iterator().(*iterator)
	if 0 != itC.h || 0 != itC.i || true != itC.forwards || nil != itC.pair {
		t.Fatal()
	}
	for x := 0; x < 10; x++ {
		if false == itC.Previous() {
			t.Fatal()
		}
		pairList = append(pairList, itC.Pair())
		valueList = append(valueList, itC.Value())
		keyList = append(keyList, itC.Key())
	}
	verifyLists()
	if 0 != itC.h || -1 != itC.i || false != itC.forwards || nil == itC.pair {
		t.Fatalf("%v", itC)
	}
	if true != itC.Next() {
		t.Fatal()
	}
	if 0 != itC.h || 1 != itC.i || true != itC.forwards || nil == itC.pair {
		t.Fatal()
	}
	if true != itC.Previous() {
		t.Fatal()
	}
	if 0 != itC.h || -1 != itC.i || false != itC.forwards || nil == itC.pair {
		t.Fatalf("%v", itC)
	}
	if false != itC.Previous() {
		t.Fatal()
	}
	if -1 != itC.h || 0 != itC.i || false != itC.forwards || nil == itC.pair {
		t.Fatalf("%v", itC)
	}
	for true == itC.Next() {
		pairList = append(pairList, itC.Pair())
		valueList = append(valueList, itC.Value())
		keyList = append(keyList, itC.Key())
	}
	verifyLists()
	for true == itC.Previous() {
		pairList = append(pairList, itC.Pair())
		valueList = append(valueList, itC.Value())
		keyList = append(keyList, itC.Key())
	}
	verifyLists()
	for true == itC.Next() {
		pairList = append(pairList, itC.Pair())
		valueList = append(valueList, itC.Value())
		keyList = append(keyList, itC.Key())
	}
	verifyLists()
	if 4 != itC.h || 0 != itC.i || true != itC.forwards || nil == itC.pair {
		t.Fatalf("%v", itC)
	}

	onePair := itC.Pair()
	// test partial backwards then partial forwards
	// 1
	if false == itC.Previous() {
		t.Fatal()
	}
	if false == comparePair(onePair, itC.Pair()) {
		t.Fatal()
	}
	// 2
	if false == itC.Previous() {
		t.Fatal()
	}
	twoPair := itC.Pair()
	// 2
	if false == itC.Next() {
		t.Fatal()
	}
	if false == comparePair(twoPair, itC.Pair()) {
		t.Fatal()
	}
	// 1
	if true != itC.Next() {
		t.Fatal()
	}
	if false == comparePair(onePair, itC.Pair()) {
		t.Fatal()
	}
	// 1
	if false != itC.Next() {
		t.Fatal()
	}
	if false == comparePair(onePair, itC.Pair()) {
		t.Fatal()
	}
	// 1
	if true == itC.Next() {
		t.Fatal()
	}
	if false == comparePair(onePair, itC.Pair()) {
		t.Fatal()
	}
	// 1
	if false == itC.Previous() {
		t.Fatal()
	}
	if false == comparePair(onePair, itC.Pair()) {
		t.Fatal()
	}
	pairList = append(pairList, itC.Pair())
	valueList = append(valueList, itC.Value())
	keyList = append(keyList, itC.Key())
	// 2
	if false == itC.Previous() {
		t.Fatal()
	}
	if false == comparePair(twoPair, itC.Pair()) {
		t.Fatal()
	}
	pairList = append(pairList, itC.Pair())
	valueList = append(valueList, itC.Value())
	keyList = append(keyList, itC.Key())
	for x := 0; x < 8; x++ {
		if false == itC.Previous() {
			t.Fatal()
		}
		pairList = append(pairList, itC.Pair())
		valueList = append(valueList, itC.Value())
		keyList = append(keyList, itC.Key())
	}
	verifyLists()

	// add some random nil values, and an empty list key
	m.m[99] = nil
	m.m[98] = []Pair{}
	m.m[0] = append(m.m[0], nil)
	m.m[1] = append(m.m[1], m.m[1][0])
	m.m[1][0] = nil
	itC = m.Iterator().(*iterator)
	for true == itC.Next() {
		pairList = append(pairList, itC.Pair())
		valueList = append(valueList, itC.Value())
		keyList = append(keyList, itC.Key())
	}
	verifyLists()

	tempM := m.m
	m.m = nil
	if true == m.Iterator().Previous() || true == m.Iterator().Next() {
		t.Fatal()
	}
	m.m = tempM

	itC.pair = nil

	if nil != itC.Value() || nil != itC.Key() {
		t.Fatal()
	}
}
