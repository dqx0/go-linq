package linq_test

import (
	"reflect"
	"testing"

	"github.com/dqx0/linq"
)

type Person struct {
	Name string
	Age  int
}

// Distinct と DistinctComparable のテスト
func TestDistinct(t *testing.T) {
	q1 := linq.New([]int{1, 2, 2, 3, 3, 3})

	expected1 := []int{1, 2, 3}
	var result1 []int
	linq.DistinctComparable(q1)(func(v int) bool {
		result1 = append(result1, v)
		return true
	})

	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("DistinctComparable failed. Expected %v, got %v", expected1, result1)
	}

	q2 := linq.New([][]int{{1, 2}, {1, 2}, {3, 4}})
	expected2 := [][]int{{1, 2}, {3, 4}}
	var result2 [][]int
	linq.Distinct(q2)(func(v []int) bool {
		result2 = append(result2, v)
		return true
	})

	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Distinct failed. Expected %v, got %v", expected2, result2)
	}

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 30}, // 重複する要素
		{Name: "Charlie", Age: 35},
	}

	q3 := linq.New(people)

	expected3 := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	var result3 []Person
	linq.Distinct(q3)(func(v Person) bool {
		result3 = append(result3, v)
		return true
	})

	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Distinct failed. Expected %v, got %v", expected3, result3)
	}
}
