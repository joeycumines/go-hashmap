# simhash
--
    import "github.com/joeycumines/go-hashmap/simhash"

Package simhash provides a simple hashmap implementation and related interfaces,
that are styled after Java's Map interface, intended to be used to store values
on keys which implement a hashing and equality function, the same as Object in
Java.

## Usage

#### type Iterator

```go
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
```

Iterator provides an interface for iterating over the key-pairs of a map. It's
value should initially start as nil, or uninitialized, and on the first call to
either Next or Previous should determine the initial direction. After being
initialized, Next (or Previous) can be used increment or de-increment the
pointer, which will always stop at the last (or first) item in the iteration.
The iterator can change direction, but will not move for the first call to the
step in the opposite direction.

#### type Key

```go
type Key interface {
	Hash() int
	Equals(other interface{}) bool
}
```

Key is an interface based on the core Java Object, which forms the basis for the
hashmap implementation. See
https://www.sitepoint.com/how-to-implement-javas-hashcode-correctly/ or look
around for details on implementing this.

#### type Map

```go
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

	// Get a new Iterator for this map, which should be stable.
	Iterator() Iterator
}
```

Map provides the specification for a hashmap type that behaves similarly to
Java's implementation, and is exported in place of the struct that implements
it.

#### func  NewMap

```go
func NewMap() Map
```

#### type Pair

```go
type Pair interface {
	Key() Key
	Value() Value
}
```

Pair represents a key-value relationship.

#### func  NewPair

```go
func NewPair(key Key, value Value) Pair
```
NewPair creates a key-value pair.

#### type Value

```go
type Value interface{}
```

Value is a value in the map, and is simply an interface{}.
