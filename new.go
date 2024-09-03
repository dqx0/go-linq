package linq

import (
	"iter"
	"slices"
)

type query[T any] iter.Seq[T]

// New はスライスをクエリに変換します。
// Converts a slice to a query.
func New[Slice ~[]T, T any](s Slice) query[T] {
	return query[T](slices.Values(s))
}
