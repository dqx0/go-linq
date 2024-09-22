package linq_test

import (
	"testing"

	"github.com/dqx0/linq"
)

func TestLast(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	r, err := linq.Last(q)
	if err != nil {
		t.Error("Expected no error but got an error")
	}
	if r != 5 {
		t.Errorf("Expected 5 but got %d", r)
	}
}

func TestLastOrDefault(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	r := linq.LastOrDefault(q, 0)
	if r != 5 {
		t.Errorf("Expected 5 but got %d", r)
	}

	list = []int{}

	q = linq.New(list)
	r = linq.LastOrDefault(q, 0)
	if r != 0 {
		t.Errorf("Expected 0 but got %d", r)
	}
}
