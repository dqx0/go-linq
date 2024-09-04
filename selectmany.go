package linq

// SelectMany は query[T] の各要素に対して指定された selector 関数を適用し、
// 得られたコレクションのすべての要素を含む新しい query[T] を返します。
//
// SelectMany applies the specified selector function to each element of the query[T],
// and returns a new query[T] containing all the elements of the collections produced by the selector.
//
// パラメータ:
//   - selector: 各要素に適用される変換関数。返り値はコレクションです。
//
// Parameters:
//   - selector: A transformation function that is applied to each element. The function returns a collection.
func SelectMany[T, U any](q Query[T], selector func(T) Query[U]) Query[U] {
	return func(yield func(U) bool) {
		for v := range q {
			innerQuery := selector(v)
			for innerValue := range innerQuery {
				if !yield(innerValue) {
					return
				}
			}
		}
	}
}
