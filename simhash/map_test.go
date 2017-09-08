package simhash

import (
	"testing"
	"sort"
)

func TestHashMap_lookup_nilKey(t *testing.T) {
	m := &hashMap{}
	h, i, ok := m.lookup(nil)
	if 0 != h || 0 != i || false != ok {
		t.Fatal()
	}
}

func TestHashMap_lookup_missingKey(t *testing.T) {
	m := &hashMap{}
	h, i, ok := m.lookup(testKeyInt(4))
	if 0 != h || 0 != i || false != ok {
		t.Fatal()
	}
	h, i, ok = m.lookup(nil)
	if 0 != h || 0 != i || false != ok {
		t.Fatal()
	}
}

func TestHashMap_lookup_noItems(t *testing.T) {
	m := &hashMap{map[int][]Pair{4: {nil}}, 0}
	h, i, ok := m.lookup(testKeyInt(4))
	if 0 != h || 0 != i || false != ok {
		t.Fatal()
	}
}

func TestHashMap_lookup(t *testing.T) {
	m := &hashMap{map[int][]Pair{4: {nil, NewPair(nil, 10), NewPair(testKeyInt(4), 20)}}, 0}
	h, i, ok := m.lookup(testKeyInt(4))
	if 4 != h || 2 != i || true != ok {
		t.Fatal()
	}
}

func TestHashMap_Size(t *testing.T) {
	m := &hashMap{nil, 55}
	if 55 != m.Size() {
		t.Fatal()
	}
}

func TestNewMap(t *testing.T) {
	m := NewMap().(*hashMap)
	if nil == m.m || 0 != len(m.m) || 0 != m.size {
		t.Fatal()
	}
}

func TestHashMap_Contains(t *testing.T) {
	m := &hashMap{map[int][]Pair{4: {nil, NewPair(nil, 10), NewPair(testKeyInt(4), 20)}}, 0}
	if true != m.Contains(testKeyInt(4)) {
		t.Fatal()
	}
	m.Put(nil, "one")
	if true != m.Contains(nil) {
		t.Fatal()
	}
}

func TestHashMap_Contains_empty(t *testing.T) {
	m := &hashMap{map[int][]Pair{4: {nil}}, 0}
	if false != m.Contains(testKeyInt(4)) && false != m.Contains(nil) {
		t.Fatal()
	}
}

func TestHashMap_Get(t *testing.T) {
	m := &hashMap{map[int][]Pair{4: {nil, NewPair(nil, 10), NewPair(testKeyInt(4), 20)}}, 0}
	if 20 != m.Get(testKeyInt(4)).(int) {
		t.Fatal()
	}
}

func TestHashMap_Get_empty(t *testing.T) {
	m := &hashMap{map[int][]Pair{4: {nil}}, 0}
	if nil != m.Get(testKeyInt(4)) {
		t.Fatal()
	}
}

func TestHashMap_Put_nilKey(t *testing.T) {
	m := NewMap().(*hashMap)
	m.size = 2
	if nil != m.Put(nil, "val") || 3 != m.size || "val" != m.Get(nil) || 1 != len(m.m) {
		t.Fatal()
	}
}

func TestHashMap_Remove_nilKey(t *testing.T) {
	m := NewMap().(*hashMap)
	if nil != m.Put(nil, "val") || 1 != m.size || "val" != m.Get(nil) || 1 != len(m.m) || true != m.Contains(nil) {
		t.Fatal()
	}
	if nil != m.Remove(testKeyInt(65)) || 1 != m.size || "val" != m.Get(nil) || 1 != len(m.m) || true != m.Contains(nil) {
		t.Fatal()
	}
	if "val" != m.Remove(nil) || 0 != m.size || nil != m.Get(nil) || 0 != len(m.m) || false != m.Contains(nil) {
		t.Fatal()
	}
}

func TestHashMap_Put_existing(t *testing.T) {
	m := &hashMap{map[int][]Pair{4: {nil, NewPair(nil, 10), NewPair(testKeyInt(4), 20)}}, 2}
	if 20 != m.Put(testKeyInt(4), 22).(int) ||
		2 != m.size ||
		22 != m.m[4][2].Value().(int) ||
		22 != m.Get(testKeyInt(4)).(int) ||
		nil != m.m[4][0] ||
		10 != m.m[4][1].Value().(int) ||
		1 != len(m.m) ||
		3 != len(m.m[4]) {
		t.Fatal()
	}
}

func TestHashMap_Put_noHash(t *testing.T) {
	m := NewMap().(*hashMap)
	if nil != m.Put(testKeyInt(4), "s") ||
		1 != len(m.m) ||
		1 != len(m.m[4]) ||
		1 != cap(m.m[4]) ||
		4 != m.m[4][0].Key().Hash() ||
		"s" != m.m[4][0].Value().(string) ||
		1 != m.size {
		t.Fatal()
	}
}

func TestHashMap_Put_nilSlice(t *testing.T) {
	m := NewMap().(*hashMap)
	m.m[4] = nil
	if nil != m.Put(testKeyInt(4), "s") ||
		1 != len(m.m) ||
		1 != len(m.m[4]) ||
		1 != cap(m.m[4]) ||
		4 != m.m[4][0].Key().Hash() ||
		"s" != m.m[4][0].Value().(string) ||
		1 != m.size {
		t.Fatal()
	}
}

func TestHashMap_Put(t *testing.T) {
	m := NewMap().(*hashMap)
	if nil != m.Put(testKeyInt(67), "67") || 1 != m.Size() || "67" != m.Get(testKeyInt(67)).(string) {
		t.Fatal()
	}
	if "67" != m.Put(testKeyInt(67), "sixty seven").(string) || 1 != m.Size() || "sixty seven" != m.Get(testKeyInt(67)).(string) {
		t.Fatal()
	}
	if nil != m.Put(testKeyInt(68), "68") || 2 != m.Size() || "68" != m.Get(testKeyInt(68)).(string) {
		t.Fatal()
	}
}

func TestHashMap_Remove(t *testing.T) {
	m := NewMap().(*hashMap)
	m.Put(testKeyInt(67), "67")
	m.Put(testKeyInt(68), "68")
	if nil != m.Remove(nil) {
		t.Fatal()
	}
	if nil != m.Remove(testKeyInt(69)) {
		t.Fatal()
	}
	if 2 != m.Size() {
		t.Fatal()
	}
	if "68" != m.Remove(testKeyInt(68)).(string) || 1 != m.Size() || false != m.Contains(testKeyInt(68)) {
		t.Fatal()
	}
	if "67" != m.Get(testKeyInt(67)).(string) || 1 != len(m.m) {
		t.Fatal()
	}
	m.m[67] = append(m.m[67], NewPair(testKeyInt(67), "sixty seven"))
	m.size++
	if "67" != m.Remove(testKeyInt(67)).(string) || 1 != m.Size() || true != m.Contains(testKeyInt(67)) {
		t.Fatalf("unexpected: %v", m.Get(testKeyInt(67)))
	}
	if "sixty seven" != m.Get(testKeyInt(67)).(string) || 1 != len(m.m) {
		t.Fatal()
	}
}

func genTestStructureHashMap() *hashMap {
	m := NewMap().(*hashMap)
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			i := x*10 + y
			m.Put(testKeyStruct{x, i}, i)
		}
	}
	m.Put(nil, 0)
	m.m[0] = append(m.m[0], nil)
	return m
}

func TestHashMap_Keys(t *testing.T) {
	m := genTestStructureHashMap()
	keys := m.Keys()
	if 10 != len(keys) {
		t.Fatalf("unexpected: %v", keys)
	}
	list := make([]int, 0)
	nilCount := 0
	for _, k := range keys {
		if nil == k {
			nilCount++
			continue
		}
		list = append(list, k.Hash())
	}
	if 1 != nilCount {
		t.Fatal()
	}
	sort.Ints(list)
	if 1 != list[0] || 1 != list[1] || 1 != list[2] ||
		2 != list[3] || 2 != list[4] || 2 != list[5] ||
		3 != list[6] || 3 != list[7] || 3 != list[8] {
		t.Fatal()
	}
}

func TestHashMap_Keys_panic(t *testing.T) {
	defer func() {
		err := recover().(error)
		if nil == err || "the output key slice was not the expected length" != err.Error() {
			t.Fatal()
		}
	}()
	m := genTestStructureHashMap()
	m.size++
	m.Keys()
}

func TestHashMap_Values_panic(t *testing.T) {
	defer func() {
		err := recover().(error)
		if nil == err || "the output value slice was not the expected length" != err.Error() {
			t.Fatal()
		}
	}()
	m := genTestStructureHashMap()
	m.size++
	m.Values()
}

func TestHashMap_Pairs_panic(t *testing.T) {
	defer func() {
		err := recover().(error)
		if nil == err || "the output pair slice was not the expected length" != err.Error() {
			t.Fatal(err.Error())
		}
	}()
	m := genTestStructureHashMap()
	m.size++
	m.Pairs()
}

func TestHashMap_Values(t *testing.T) {
	m := genTestStructureHashMap()
	values := m.Values()
	list := make([]int, 0)
	for _, v := range values {
		list = append(list, v.(int))
	}
	if 10 != len(list) {
		t.Fatalf("unexpected: %v", list)
	}
	sort.Ints(list)
	if 0 != list[0] ||
		11 != list[1] || 12 != list[2] || 13 != list[3] ||
		21 != list[4] || 22 != list[5] || 23 != list[6] ||
		31 != list[7] || 32 != list[8] || 33 != list[9] {
		t.Fatalf("unexpected: %v", list)
	}
}

func TestHashMap_Pairs(t *testing.T) {
	m := genTestStructureHashMap()
	pairs := m.Pairs()
	if 10 != len(pairs) {
		t.Fatalf("unexpected: %v", pairs)
	}
	keyList := make([]int, 0)
	valueList := make([]int, 0)
	nilCount := 0
	for _, p := range pairs {
		if nil == p {
			t.Fatal("nil pair")
		}
		valueList = append(valueList, p.Value().(int))
		if nil == p.Key() {
			nilCount++
			continue
		}
		keyList = append(keyList, p.Key().Hash())
	}
	if 1 != nilCount {
		t.Fatal()
	}
	sort.Ints(keyList)
	if 1 != keyList[0] || 1 != keyList[1] || 1 != keyList[2] ||
		2 != keyList[3] || 2 != keyList[4] || 2 != keyList[5] ||
		3 != keyList[6] || 3 != keyList[7] || 3 != keyList[8] {
		t.Fatal()
	}
	if 10 != len(valueList) {
		t.Fatalf("unexpected: %v", valueList)
	}
	sort.Ints(valueList)
	sort.Ints(valueList)
	if 0 != valueList[0] ||
		11 != valueList[1] || 12 != valueList[2] || 13 != valueList[3] ||
		21 != valueList[4] || 22 != valueList[5] || 23 != valueList[6] ||
		31 != valueList[7] || 32 != valueList[8] || 33 != valueList[9] {
		t.Fatalf("unexpected: %v", valueList)
	}
}

func TestHashMap_Serialize(t *testing.T) {
	m := genTestStructureHashMap()
	s := m.Serialize()
	if nil == s {
		t.Fatal()
	}
	if 10 != len(s) {
		t.Fatal()
	}
	if 0 != s["<nil>"].(int) {
		t.Fatal()
	}
	if 11 != s["11"].(int) {
		t.Fatal()
	}
	if 12 != s["12"].(int) {
		t.Fatal()
	}
	if 13 != s["13"].(int) {
		t.Fatal()
	}
	if 21 != s["21"].(int) {
		t.Fatal()
	}
	if 22 != s["22"].(int) {
		t.Fatal()
	}
	if 23 != s["23"].(int) {
		t.Fatal()
	}
	if 31 != s["31"].(int) {
		t.Fatal()
	}
	if 32 != s["32"].(int) {
		t.Fatal()
	}
	if 33 != s["33"].(int) {
		t.Fatal()
	}
}

func TestHashMap_Iterator(t *testing.T) {
	m := genTestStructureHashMap()
	pairs := make([]Pair, 0)
	for it := m.Iterator(); true == it.Next(); {
		//fmt.Printf("%v\n", pair)
		pairs = append(pairs, it.Pair())
	}
	if 10 != len(pairs) {
		t.Fatalf("unexpected: %v", pairs)
	}
	keyList := make([]int, 0)
	valueList := make([]int, 0)
	nilCount := 0
	for _, p := range pairs {
		if nil == p {
			t.Fatal("nil pair")
		}
		valueList = append(valueList, p.Value().(int))
		if nil == p.Key() {
			nilCount++
			continue
		}
		keyList = append(keyList, p.Key().Hash())
	}
	if 1 != nilCount {
		t.Fatal()
	}
	sort.Ints(keyList)
	if 1 != keyList[0] || 1 != keyList[1] || 1 != keyList[2] ||
		2 != keyList[3] || 2 != keyList[4] || 2 != keyList[5] ||
		3 != keyList[6] || 3 != keyList[7] || 3 != keyList[8] {
		t.Fatal()
	}
	if 10 != len(valueList) {
		t.Fatalf("unexpected: %v", valueList)
	}
	sort.Ints(valueList)
	sort.Ints(valueList)
	if 0 != valueList[0] ||
		11 != valueList[1] || 12 != valueList[2] || 13 != valueList[3] ||
		21 != valueList[4] || 22 != valueList[5] || 23 != valueList[6] ||
		31 != valueList[7] || 32 != valueList[8] || 33 != valueList[9] {
		t.Fatalf("unexpected: %v", valueList)
	}
}
