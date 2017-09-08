package simhash

import "testing"

func TestPair(t *testing.T) {
	k := testKeyInt(4)
	p := NewPair(k, "test")
	if nil == p {
		t.Fatalf("unexpected: %v", p)
	}
	if 4 != p.Key().Hash() {
		t.Fatalf("unexpected: %v", p)
	}
	if "test" != p.Value().(string) {
		t.Fatalf("unexpected: %v", p)
	}
}
