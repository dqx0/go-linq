package linq_test

import (
	"testing"

	"github.com/dqx0/linq"
)

type Inner struct {
	Value int
}

type Outer struct {
	Inners []Inner
}

func TestSelectManySimpleStruct(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	s := linq.SelectMany(q, func(item int) linq.Query[int] {
		return linq.New([]int{item, item * 2})
	}).ToSlice()

	expected := []int{1, 2, 2, 4, 3, 6, 4, 8, 5, 10}

	if len(s) != len(expected) {
		t.Errorf("Expected length %d but got %d", len(expected), len(s))
	}

	for i, v := range s {
		if v != expected[i] {
			t.Errorf("Expected %d but got %d", expected[i], v)
			break
		}
	}
}
func TestSelectManyNestedStruct(t *testing.T) {
	list := []Outer{
		{Inners: []Inner{{Value: 1}, {Value: 2}}},
		{Inners: []Inner{{Value: 3}, {Value: 4}}},
		{Inners: []Inner{{Value: 5}}},
	}

	q := linq.New(list)
	s := linq.SelectMany(q, func(item Outer) linq.Query[Inner] {
		return linq.New(item.Inners)
	}).ToSlice()

	expected := []Inner{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}}

	if len(s) != len(expected) {
		t.Errorf("Expected length %d but got %d", len(expected), len(s))
	}

	for i, v := range s {
		if v != expected[i] {
			t.Errorf("Expected %v but got %v", expected[i], v)
			break
		}
	}
}
