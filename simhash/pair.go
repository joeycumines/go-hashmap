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
