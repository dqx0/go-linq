package linq

// Select は query[T] の各要素に対して指定された selector 関数を適用し、
// 新しい query[T] を返します。
//
// Select applies the specified selector function to each element of the query[T],
// and returns a new query[T].
// The parameter 'a' takes additional string information but does not affect the method's behavior.
//
// パラメータ:
//   - selector: 各要素に適用される変換関数。
//
// Parameters:
//   - selector: A transformation function to apply to each element.
func Select[T any, U any](q Query[T], selector func(item T) U) Query[U] {
	return func(yield func(U) bool) {
		for v := range q {
			if !yield(selector(v)) {
				return
			}
		}
	}
}
