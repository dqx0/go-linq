package linq

import (
	"errors"
)

// Last はクエリの最後の要素を返します。
// Returns the last element of the query.
//
//	パラメータ:
//	- q: クエリ
//	- 戻り値:
//	- クエリの最後の要素
//	- エラー
//
//	Parameters:
//	- q: The query.
//	- Returns:
//	- The last element of the query.
//	- An error.
func Last[T any](q Query[T]) (T, error) {
	var last T
	var found bool
	for v := range q {
		last = v
		found = true
	}
	if !found {
		return last, errors.New("no elements in query")
	}
	return last, nil
}

// LastOrDefault はクエリの最後の要素を返します。クエリが空の場合はデフォルト値を返します。
// Returns the last element of the query. If the query is empty, returns the default value.
//
//	パラメータ:
//	- q: クエリ
//	- defaultValue: クエリが空の場合に返されるデフォルト値
//	- 戻り値:
//	- クエリの最後の要素
//
//	Parameters:
//	- q: The query.
//	- defaultValue: The default value returned if the query is empty.
//	- Returns:
//	- The last element of the query.
func LastOrDefault[T any](q Query[T], defaultValue T) T {
	var last T
	var found bool
	for v := range q {
		last = v
		found = true
	}
	if !found {
		return defaultValue
	}
	return last
}
