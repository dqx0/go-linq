package linq_test

import (
	"testing"

	"github.com/dqx0/linq"
)

type Wrapper1 struct {
	Value1 string
	Value2 int
}
type Wrapper2 struct {
	Wrapper1 Wrapper1
}

func TestContainsSimple(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	if !linq.Contains(q, 3) {
		t.Error("Expected true but got false")
	}
	if linq.Contains(q, 6) {
		t.Error("Expected false but got true")
	}
}

func TestContainsComplex(t *testing.T) {
	list := []Wrapper2{
		{Wrapper1{Value1: "1", Value2: 1}},
		{Wrapper1{Value1: "2", Value2: 2}},
		{Wrapper1{Value1: "3", Value2: 3}},
		{Wrapper1{Value1: "4", Value2: 4}},
		{Wrapper1{Value1: "5", Value2: 5}},
	}

	q := linq.New(list)
	if !linq.Contains(q, Wrapper2{Wrapper1{Value1: "3", Value2: 3}}) {
		t.Error("Expected true but got false")
	}
	if linq.Contains(q, Wrapper2{Wrapper1{Value1: "6", Value2: 6}}) {
		t.Error("Expected false but got true")
	}
}
