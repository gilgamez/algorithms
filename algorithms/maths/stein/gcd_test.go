package gcd

import (
	"testing"
)

func TestRecurse(t *testing.T) {
	if Recurse(14, 7) != 7 ||
		Recurse(4, 2) != 2 ||
		Recurse(31, 2) != 1 ||
		Recurse(33, 11) != 11 {
		t.Error()
	}
}

func TestIterative(t *testing.T) {
	if Iter(14, 7) != 7 ||
		Iter(4, 2) != 2 ||
		Iter(31, 2) != 1 ||
		Iter(33, 11) != 11 {
		t.Error()
	}
}

func BenchmarkRecurse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recurse(131313131, 12343545)
	}
}

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iter(131313131, 12343545)
	}
}
