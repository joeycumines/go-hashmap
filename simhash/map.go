// Package simhash provides a simple hashmap implementation and related interfaces, that are styled after Java's Map
// interface, intended to be used to store values on keys which implement a hashing and equality function, the same
// as Object in Java.
package simhash

import (
	"errors"
	"fmt"
)

// Map provides the specification for a hashmap type that behaves similarly to Java's implementation, and is exported
// in place of the struct that implements it.
type Map interface {
	// Contains will return true if the key exists in the map.
	Contains(key Key) bool

	// Get will return the value if it exists in the map, or nil if it doesn't.
	Get(key Key) Value

	// Put will store value as key in the map, and will return any existing value, or nil.
	Put(key Key, value Value) Value

	// Remove removes any value that existed for key in the map, and will return it, or nil.
	Remove(key Key) Value

	// Keys returns a slice containing all the keys in the map.
	Keys() []Key

	// Values returns a slice containing all the values in the map.
	Values() []Value

	// Pairs returns a slice containing all the key-value pairs in the map.
	Pairs() []Pair

	// Size returns the number of key-value pairs in the map.
	Size() int

	// Serialize should a best-fit string serialization the keys in the map, associated with their values. This method
	// is not guaranteed to return a map of the same size, and is dependant on the key implementation - if the keys
	// all implement fmt.Stringer, and return correctly serialized values, this will work correctly.
	Serialize() map[string]interface{}
}

// A similar implementation to the HashMap in Java, this uses the underlying Go map but allows efficient (citation
// needed) lookup of structs which either cannot be easily serialized or are expensive to do so.
// Will most certainly break under unsynchronised concurrent write conditions, concurrent reads will be ok if the
// map isn't changing (and the state is assured to be complete).
type hashMap struct {
	m    map[int][]Pair
	size int
}

func (m *hashMap) lookup(key Key) (int, int, bool) {
	h := 0
	if nil != key {
		h = key.Hash()
	}
	pairs, ok := m.m[h]
	if false == ok || nil == pairs {
		return 0, 0, false
	}
	for i, pair := range pairs {
		if nil == pair {
			continue
		}
		if k := pair.Key(); (nil == key && nil == k) || (nil != key && nil != k && key.Equals(k)) {
			return h, i, true
		}
	}
	return 0, 0, false
}

func (m *hashMap) Contains(key Key) bool {
	_, _, ok := m.lookup(key)
	return ok
}

func (m *hashMap) Get(key Key) Value {
	h, i, ok := m.lookup(key)
	if false == ok {
		return nil
	}
	return m.m[h][i].Value()
}

func (m *hashMap) Put(key Key, value Value) Value {
	if h, i, ok := m.lookup(key); true == ok {
		v := m.m[h][i].Value()
		m.m[h][i] = NewPair(key, value)
		return v
	}
	h := 0
	if nil != key {
		h = key.Hash()
	}
	if pairs, ok := m.m[h]; false == ok || nil == pairs {
		m.m[h] = make([]Pair, 0, 1)
	}
	m.m[h] = append(m.m[h], NewPair(key, value))
	m.size++
	return nil
}

func (m *hashMap) Remove(key Key) Value {
	h, i, ok := m.lookup(key)
	if false == ok {
		return nil
	}
	v := m.m[h][i].Value()
	m.m[h][i] = m.m[h][len(m.m[h])-1]
	m.m[h][len(m.m[h])-1] = nil
	m.m[h] = m.m[h][:len(m.m[h])-1]
	if 0 == len(m.m[h]) {
		delete(m.m, h)
	}
	m.size--
	return v
}

func (m *hashMap) Keys() []Key {
	keys := make([]Key, m.size)
	x := 0
	for _, pairs := range m.m {
		for _, pair := range pairs {
			if nil == pair {
				continue
			}
			k := pair.Key()
			keys[x] = k
			x++
		}
	}
	if x != m.size {
		panic(errors.New("the output key slice was not the expected length"))
	}
	return keys
}

func (m *hashMap) Values() []Value {
	values := make([]Value, m.size)
	x := 0
	for _, pairs := range m.m {
		for _, pair := range pairs {
			if nil == pair {
				continue
			}
			values[x] = pair.Value()
			x++
		}
	}
	if x != m.size {
		panic(errors.New("the output value slice was not the expected length"))
	}
	return values
}

func (m *hashMap) Pairs() []Pair {
	pairList := make([]Pair, m.size)
	x := 0
	for _, pairs := range m.m {
		for _, pair := range pairs {
			if nil == pair {
				continue
			}
			pairList[x] = pair
			x++
		}
	}
	if x != m.size {
		panic(errors.New("the output pair slice was not the expected length"))
	}
	return pairList
}

func (m *hashMap) Size() int {
	return m.size
}

func (m *hashMap) Serialize() map[string]interface{} {
	serialized := make(map[string]interface{})
	for _, pairs := range m.m {
		for _, pair := range pairs {
			if nil == pair {
				continue
			}
			kStringer, ok := pair.Key().(fmt.Stringer)
			var k string
			if true == ok {
				k = kStringer.String()
			} else {
				k = fmt.Sprintf("%v", pair.Key())
			}
			serialized[k] = pair.Value()
		}
	}
	return serialized
}

func (m *hashMap) Iterator() Iterator {
	// enumerate the map keys ahead of time, they are only integers anyway
	hList := make([]int, 0, len(m.m))
	for h := range m.m {
		hList = append(hList, h)
	}
	return &iterator{
		m,
		hList,
		nil,
		0,
		0,
		true,
		false,
	}
}

func NewMap() Map {
	return &hashMap{make(map[int][]Pair), 0}
}
