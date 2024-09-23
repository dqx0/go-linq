package linq

// Count はクエリの要素数を返します。
// Returns the number of elements in the query.
//
//	パラメータ:
//	- q: クエリ
//	- 戻り値:
//	- クエリの要素数
//
//	Parameters:
//	- q: The query.
//	- Returns:
//	- The number of elements in the query.
func Count[T any](q Query[T]) int {
	count := 0
	for range q {
		count++
	}
	return count
}
