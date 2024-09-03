package linq

// スライスが空でないかどうかを返します。
// 空でない場合は true を返し、空の場合は false を返します。
//
// Returns true if the slice is not empty, otherwise returns false.
func (q query[T]) Any() bool {
	for _ = range q {
		return true
	}
	return false
}
