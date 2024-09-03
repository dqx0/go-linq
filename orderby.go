package linq

import (
	"slices"
)

func (q query[T]) OrderBy(cmp func(a, b T) int) query[T] {
	values := q.ToSlice()

	slices.SortFunc(values, cmp)

	return func(yield func(T) bool) {
		for _, v := range values {
			if !yield(v) {
				return
			}
		}
	}
}

func (q query[T]) OrderByDescending(cmp func(a, b T) int) query[T] {
	values := q.ToSlice()

	slices.SortFunc(values, cmp)
	slices.Reverse(values)

	return func(yield func(T) bool) {
		for _, v := range values {
			if !yield(v) {
				return
			}
		}
	}
}
