package linq

import (
	"iter"
	"slices"
)

type Query[T any] iter.Seq[T]

// New はスライスをクエリに変換します。
// Converts a slice to a query.
func New[Slice ~[]T, T any](s Slice) Query[T] {
	return Query[T](slices.Values(s))
}
