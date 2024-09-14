package linq_test

import (
	"testing"

	"github.com/dqx0/linq"
)

func TestFirst(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	r, err := linq.First(q)
	if err != nil {
		t.Error("Expected no error but got an error")
	}
	if r != 1 {
		t.Errorf("Expected 1 but got %d", r)
	}
}

func TestFirstOrDefault(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	r := linq.FirstOrDefault(q, 0)
	if r != 1 {
		t.Errorf("Expected 1 but got %d", r)
	}

	list = []int{}

	q = linq.New(list)
	r = linq.FirstOrDefault(q, 0)
	if r != 0 {
		t.Errorf("Expected 0 but got %d", r)
	}
}
