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

func TestDistinctBy(t *testing.T) {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 35},
		{Name: "Charlie", Age: 35},
	}

	q := linq.New(people)

	expected := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	var result []Person
	linq.DistinctBy(q, func(p Person) string {
		return p.Name
	})(func(v Person) bool {
		result = append(result, v)
		return true
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DistinctBy failed. Expected %v, got %v", expected, result)
	}
}

func TestDistinctByComparable(t *testing.T) {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 35},
		{Name: "Charlie", Age: 35},
	}

	q := linq.New(people)

	expected := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	var result []Person
	linq.DistinctByComparable(q, func(p Person) string {
		return p.Name
	})(func(v Person) bool {
		result = append(result, v)
		return true
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DistinctByComparable failed. Expected %v, got %v", expected, result)
	}
}

// DistinctとDistinctComparableのベンチマーク
func BenchmarkDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := linq.New([]int{1, 2, 2, 3, 3, 3})
		linq.Distinct(q)(func(v int) bool {
			return true
		})
	}
}

func BenchmarkDistinctComparable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := linq.New([]int{1, 2, 2, 3, 3, 3})
		linq.DistinctComparable(q)(func(v int) bool {
			return true
		})
	}
}

// DistinctByとDistinctByComparableのベンチマーク
func BenchmarkDistinctBy(b *testing.B) {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 35},
		{Name: "Charlie", Age: 35},
	}

	q := linq.New(people)

	for i := 0; i < b.N; i++ {
		linq.DistinctBy(q, func(p Person) string {
			return p.Name
		})(func(v Person) bool {
			return true
		})
	}
}

func BenchmarkDistinctByComparable(b *testing.B) {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 35},
		{Name: "Charlie", Age: 35},
	}

	q := linq.New(people)

	for i := 0; i < b.N; i++ {
		linq.DistinctByComparable(q, func(p Person) string {
			return p.Name
		})(func(v Person) bool {
			return true
		})
	}
}
