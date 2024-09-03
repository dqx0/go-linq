package linq

// Aggregate は query[T] の各要素に対して指定された aggregator 関数を適用し、
// 結果を返します。
//
// Aggregate applies the specified aggregator function to each element of the query[T],
// and returns the result.
//
// パラメータ:
//   - initial: 初期値。
//   - aggregator: 要素を集計する関数。
//
// Parameters:
//   - initial: The initial value.
//   - aggregator: A function to aggregate the elements.
func (q query[T]) Aggregate(initial T, aggregator func(accumulator T, item T) T) T {
	acc := initial
	q(func(item T) bool {
		acc = aggregator(acc, item)
		return true
	})
	return acc
}
