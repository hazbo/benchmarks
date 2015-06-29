package simple

import (
	"testing"
)

func TestPrepend(t *testing.T) {
	e := []int{2, 4, 6, 8, 10, 12, 14, 16}
	s := []int{10, 12, 14, 16}
	s  = Prepend(s, 2, 4, 6, 8)
	if len(s) != len(e) {
		t.Error("Expected slice len to be 8, got", len(e))
	}
	for i := range s {
		if s[i] != e[i] {
			t.Error("Not expecting", s[i])
		}
	}	
}

func BenchmarkPrepend(b *testing.B) {
	s := []int{10, 12, 14, 16}
    for n := 0; n < b.N; n++ {
    	s = Prepend(s, 2, 4, 6, 8)
    }
}

func TestAppend(t *testing.T) {
	e := []int{2, 4, 6, 8, 10, 12, 14, 16}
	s := []int{2, 4, 6, 8}
	s  = Append(s, 10, 12, 14, 16)
	if len(s) != len(e) {
		t.Error("Expected slice len to be 8, got", len(e))
	}
	for i := range s {
		if s[i] != e[i] {
			t.Error("Not expecting", s[i])
		}
	}
}

func BenchmarkAppend(b *testing.B) {
	s := []int{2, 4, 6, 8}
    for n := 0; n < b.N; n++ {
    	s = Append(s, 10, 12, 14, 16)
    }
}
