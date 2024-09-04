package linq

// スライスが空でないかどうかを返します。
// 空でない場合は true を返し、空の場合は false を返します。
//
// Returns true if the slice is not empty, otherwise returns false.
//
// 	パラメータ:
// 	- q: クエリ
//
// 	戻り値:
// 	- クエリが空でない場合は true、空の場合は false
//
// 	Parameters:
// 	- q: The query.
//
// 	Returns:
// 	- true if the query is not empty, otherwise false.
//
func (q Query[T]) Any() bool {
	for _ = range q {
		return true
	}
	return false
}

// Any はクエリが空でないかどうかを返します。
// 空でない場合は true を返し、空の場合は false を返します。
//
// Any returns true if the query is not empty, otherwise returns false.
//
// 	パラメータ:
// 	- q: クエリ
//
// 	戻り値:
// 	- クエリが空でない場合は true、空の場合は false
//
// 	Parameters:
// 	- q: The query.
//
// 	Returns:
// 	- true if the query is not empty, otherwise false.
//
func Any[T any](q Query[T]) bool {
	return q.Any()
}
