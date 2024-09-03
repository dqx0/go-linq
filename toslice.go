package linq

// ToSlice は query[T] をスライスに変換します。
//
// Converts the query[T] to a slice.
func (q query[T]) ToSlice() []T {
	l := make([]T, 0)
	for v := range q {
		l = append(l, v)
	}
	return l
}
