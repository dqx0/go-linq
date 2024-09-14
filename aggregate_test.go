package linq_test

import (
	"fmt"
	"testing"

	"github.com/dqx0/linq"
)

func TestAggregate(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	r := linq.Aggregate(q, "", func(acc string, item int) string {
		return acc + fmt.Sprintf("%d", item)
	})
	if r != "12345" {
		t.Errorf("Expected 15 but got %s", r)
	}
}

func TestAggregateAverage(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	total := linq.Aggregate(q, struct{ sum, count int }{0, 0}, func(acc struct{ sum, count int }, item int) struct{ sum, count int } {
		acc.sum += item
		acc.count++
		return acc
	})
	average := float64(total.sum) / float64(total.count)
	if average != 3.0 {
		t.Errorf("Expected 3.0 but got %f", average)
	}
}

func TestAggregateWithoutInitialSum(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	sum, err := linq.AggregateWithoutInitial(q, func(acc, item int) int {
		return acc + item
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sum != 15 {
		t.Errorf("Expected 15 but got %d", sum)
	}
}
func TestAggregateWithoutInitialMax(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	q := linq.New(list)
	max, err := linq.AggregateWithoutInitial(q, func(acc, item int) int {
		if item > acc {
			return item
		}
		return acc
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if max != 5 {
		t.Errorf("Expected 5 but got %d", max)
	}
}
func TestAggregateWithoutInitialEmpty(t *testing.T) {
	var list []int

	q := linq.New(list)
	_, err := linq.AggregateWithoutInitial(q, func(acc, item int) int {
		return acc + item
	})
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}
