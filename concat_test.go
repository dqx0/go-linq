package linq_test

import (
	"reflect"
	"testing"

	"github.com/dqx0/linq"
)

func TestConcat(t *testing.T) {
	q1 := linq.New([]int{1, 2, 3})
	q2 := linq.New([]int{4, 5, 6})

	expected := []int{1, 2, 3, 4, 5, 6}
	result := linq.Concat(q1, q2).ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Concat failed. Expected %v, got %v", expected, result)
	}
}

func TestAppend(t *testing.T) {
	t.Run("Append int item", func(t *testing.T) {
		q := linq.New([]int{1, 2, 3})

		newQuery := linq.Append(q, 4)

		var result []int
		newQuery(func(v int) bool {
			result = append(result, v)
			return true
		})

		expected := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Append failed: expected %v, got %v", expected, result)
		}
	})

	t.Run("Append string item", func(t *testing.T) {
		q := linq.New([]string{"apple", "banana"})

		newQuery := linq.Append(q, "cherry")

		var result []string
		newQuery(func(v string) bool {
			result = append(result, v)
			return true
		})

		expected := []string{"apple", "banana", "cherry"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Append failed: expected %v, got %v", expected, result)
		}
	})
}

func TestPrepend(t *testing.T) {
	q := linq.New([]int{1, 2, 3})

	expected := []int{4, 1, 2, 3}
	result := linq.Prepend(q, 4).ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Prepend failed. Expected %v, got %v", expected, result)
	}
}

func TestConcatEmpty(t *testing.T) {
	q1 := linq.New([]int{1, 2, 3})
	q2 := linq.New([]int{})

	expected := []int{1, 2, 3}
	result := linq.Concat(q1, q2).ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Concat failed. Expected %v, got %v", expected, result)
	}
}

func TestAppendEmpty(t *testing.T) {
	q := linq.New([]int{})

	expected := []int{1}
	result := linq.Append(q, 1).ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Append failed. Expected %v, got %v", expected, result)
	}
}

func TestPrependEmpty(t *testing.T) {
	q := linq.New([]int{})

	expected := []int{4}
	result := linq.Prepend(q, 4).ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Prepend failed. Expected %v, got %v", expected, result)
	}
}

func TestConcatEmptyBoth(t *testing.T) {
	q1 := linq.New([]int{})
	q2 := linq.New([]int{})

	expected := []int{}
	result := linq.Concat(q1, q2).ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Concat failed. Expected %v, got %v", expected, result)
	}
}

func BenchmarkConcat(b *testing.B) {
	q1 := linq.New([]int{1, 2, 3})
	q2 := linq.New([]int{4, 5, 6})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linq.Concat(q1, q2)
	}
}

func BenchmarkAppend(b *testing.B) {
	q := linq.New([]int{1, 2, 3})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linq.Append(q, 4)
	}
}

func BenchmarkPrepend(b *testing.B) {
	q := linq.New([]int{1, 2, 3})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linq.Prepend(q, 4)
	}
}
