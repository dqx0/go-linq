package linq

// Where は条件に合致する要素だけを返すクエリを返します。
//
// Returns a query that returns only the elements that match the condition.
//
// 	パラメータ:
// 	- cond: 条件
//
// 	戻り値:
// 	- 条件に合致する要素だけを返すクエリ
//
// 	Parameters:
// 	- cond: The condition.
//
// 	Returns:
// 	- A query that returns only the elements that match the condition.
//
func (q query[T]) Where(cond func(x T) bool) query[T] {
	return func(yield func(T) bool) {
		for v := range q {
			if !cond(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}
