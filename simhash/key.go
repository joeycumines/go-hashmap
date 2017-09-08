package simhash

// Key is an interface based on the core Java Object, which forms the basis for the hashmap implementation.
// See https://www.sitepoint.com/how-to-implement-javas-hashcode-correctly/ or look around for details on implementing
// this.
type Key interface {
	Hash() int
	Equals(other interface{}) bool
}
