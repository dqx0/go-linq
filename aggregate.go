package linq

import "errors"

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
func Aggregate[T, U any](q Query[T], initial U, aggregator func(accumulator U, item T) U) U {
	acc := initial
	q(func(item T) bool {
		acc = aggregator(acc, item)
		return true
	})
	return acc
}

// AggregateWithoutInitial は初期値なしで query[T] の各要素に対して指定された aggregator 関数を適用し、
// 結果を返します。最初の要素を初期値として使用します。
//
// AggregateWithoutInitial applies the specified aggregator function to each element of the query[T],
// and returns the result. The first element is used as the initial value.
//
// パラメータ:
//   - aggregator: 要素を集計する関数。
//
// Parameters:
//   - aggregator: A function to aggregate the elements.
func AggregateWithoutInitial[T any](q Query[T], aggregator func(accumulator T, item T) T) (T, error) {
	var initial T
	first := true
	var err error

	q(func(item T) bool {
		if first {
			initial = item
			first = false
		} else {
			initial = aggregator(initial, item)
		}
		return true
	})

	if first {
		err = errors.New("sequence contains no elements")
	}

	return initial, err
}
