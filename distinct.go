package linq

import (
	"reflect"
)

// Distinct はクエリから重複を取り除いた新しいクエリを返します。
// Returns a new query with duplicates removed from the query.
//
//	パラメータ:
//	- q: クエリ
//	- 戻り値:
//	- 重複を取り除いた新しいクエリ
//
//	Parameters:
//	- q: The query.
//	- Returns:
//	- A new query with duplicates removed.
//
// Distinct はクエリから重複を取り除いた新しいクエリを返します。
// Returns a new query with duplicates removed from the query.
//
//	パラメータ:
//	- q: クエリ
//	- 戻り値:
//	- 重複を取り除いた新しいクエリ
//
//	Parameters:
//	- q: The query.
//	- Returns:
//	- A new query with duplicates removed.
func Distinct[T any](q Query[T]) Query[T] {
	return func(yield func(T) bool) {
		seen := []T{}
		for v := range q {
			if !contains(seen, v) {
				seen = append(seen, v)
				if !yield(v) {
					return
				}
			}
		}
	}
}

// DistinctComparable は、T が comparable な場合、
// マップを使った効率的な重複除去を行います。
//
// DistinctComparable is optimized for comparable types, using a map for deduplication.
func DistinctComparable[T comparable](q Query[T]) Query[T] {
	return func(yield func(T) bool) {
		seen := make(map[T]struct{})
		for v := range q {
			if _, exists := seen[v]; !exists {
				seen[v] = struct{}{}
				if !yield(v) {
					return
				}
			}
		}
	}
}

func contains[T any](seen []T, value T) bool {
	for _, v := range seen {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// DistinctBy は、指定されたキーによってクエリから重複を取り除いた新しいクエリを返します。
// Returns a new query with duplicates removed from the query based on the specified key.
//
//	パラメータ:
//	- q: クエリ
//	- keySelector: キーを取得する関数
//	- 戻り値:
//	- 重複を取り除いた新しいクエリ
//
//	Parameters:
//	- q: The query.
//	- keySelector: A function to extract the key.
//	- Returns:
//	- A new query with duplicates removed.
func DistinctBy[T any, K any](q Query[T], keySelector func(T) K) Query[T] {
	return func(yield func(T) bool) {
		seen := []K{} // 見つけたキーを保持するスライス
		for v := range q {
			key := keySelector(v)
			found := false
			for _, s := range seen {
				if reflect.DeepEqual(s, key) {
					found = true
					break
				}
			}
			if !found {
				seen = append(seen, key)
				if !yield(v) {
					return
				}
			}
		}
	}
}

// DistinctByComparable は、T が comparable な場合、
// 効率的な重複除去を行います。
//
// DistinctByComparable is optimized for comparable types, using a map for deduplication.
//
//	パラメータ:
//	- q: クエリ
//	- keySelector: キーを取得する関数
//	- 戻り値:
//	- 重複を取り除いた新しいクエリ
//
//	Parameters:
//	- q: The query.
//	- keySelector: A function to extract the key.
//	- Returns:
//	- A new query with duplicates removed.
func DistinctByComparable[T any, K comparable](q Query[T], keySelector func(T) K) Query[T] {
	return func(yield func(T) bool) {
		seen := make(map[K]struct{})
		for v := range q {
			key := keySelector(v)
			if _, exists := seen[key]; !exists {
				seen[key] = struct{}{}
				if !yield(v) {
					return
				}
			}
		}
	}
}
