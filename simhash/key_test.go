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
