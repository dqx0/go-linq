package linq

import (
	"slices"
)

// OrderBy はクエリの要素を指定された比較関数でソートします。
//
// OrderBy sorts the elements of the query using the specified comparison function.
//
// パラメータ:
//   - q: ソートするクエリ
//   - cmp: 比較関数
//
// Parameters:
//   - q: The query to sort.
//   - cmp: The comparison function.
func OrderBy[T any](q Query[T], cmp func(a, b T) int) Query[T] {
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

// OrderByDescending はクエリの要素を指定された比較関数で降順にソートします。
//
// OrderByDescending sorts the elements of the query in descending order using the specified comparison function.
//
// パラメータ:
//   - q: ソートするクエリ
//   - cmp: 比較関数
//
// Parameters:
//   - q: The query to sort.
//   - cmp: The comparison function.
func OrderByDescending[T any](q Query[T], cmp func(a, b T) int) Query[T] {
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
