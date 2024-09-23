package linq_test

import (
	"testing"

	"github.com/dqx0/linq"
)

func TestCount(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	r := linq.Count(q)
	if r != 5 {
		t.Errorf("Expected 5 but got %d", r)
	}
}

func TestCountEmpty(t *testing.T) {
	list := []int{}

	q := linq.New(list)
	r := linq.Count(q)
	if r != 0 {
		t.Errorf("Expected 0 but got %d", r)
	}
}
