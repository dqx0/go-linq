package linq

import (
	"errors"
)

// First はクエリの最初の要素を返します。
// Returns the first element of the query.
//
//	パラメータ:
//	- q: クエリ
//	- 戻り値:
//	- クエリの最初の要素
//	- エラー
//
//	Parameters:
//	- q: The query.
//	- Returns:
//	- The first element of the query.
//	- An error.
func First[T any](q Query[T]) (T, error) {
	for v := range q {
		return v, nil
	}
	var zero T
	return zero, errors.New("no elements in query")
}

// FirstOrDefault はクエリの最初の要素を返します。クエリが空の場合はデフォルト値を返します。
// Returns the first element of the query. If the query is empty, returns the default value.
//
//	パラメータ:
//	- q: クエリ
//	- defaultValue: クエリが空の場合に返されるデフォルト値
//	- 戻り値:
//	- クエリの最初の要素
//
//	Parameters:
//	- q: The query.
//	- defaultValue: The default value returned if the query is empty.
//	- Returns:
//	- The first element of the query.
func FirstOrDefault[T any](q Query[T], defaultValue T) T {
	for v := range q {
		return v
	}
	return defaultValue
}
