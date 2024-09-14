package linq_test

import (
	"fmt"
	"testing"

	"github.com/dqx0/linq"
)

func TestSelect(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	r := linq.Select(q, func(v int) int {
		return v * 2
	})
	l := r.ToSlice()
	if len(l) != 5 {
		t.Errorf("Expected 5 but got %d", len(l))
	}
	if l[0] != 2 {
		t.Errorf("Expected 2 but got %d", l[0])
	}
	if l[1] != 4 {
		t.Errorf("Expected 4 but got %d", l[1])
	}
	if l[2] != 6 {
		t.Errorf("Expected 6 but got %d", l[2])
	}
	if l[3] != 8 {
		t.Errorf("Expected 8 but got %d", l[3])
	}
	if l[4] != 10 {
		t.Errorf("Expected 10 but got %d", l[4])
	}
}

func TestSelectDifferentTypes(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	r := linq.Select(q, func(v int) string {
		return fmt.Sprintf("Number: %d", v)
	})
	l := r.ToSlice()
	expected := []string{"Number: 1", "Number: 2", "Number: 3", "Number: 4", "Number: 5"}
	if len(l) != len(expected) {
		t.Errorf("Expected length %d but got %d", len(expected), len(l))
	}
	for i, v := range l {
		if v != expected[i] {
			t.Errorf("Expected %s but got %s", expected[i], v)
		}
	}
}
