package linq

// ToSlice は query[T] をスライスに変換します。
//
// Converts the query[T] to a slice.
func (q Query[T]) ToSlice() []T {
	l := make([]T, 0)
	for v := range q {
		l = append(l, v)
	}
	return l
}

// ToSlice は query[T] をスライスに変換します。
//
// Converts the query[T] to a slice.
func ToSlice[T any](q Query[T]) []T {
	return q.ToSlice()
}
