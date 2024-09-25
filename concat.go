package linq

// Apped は、クエリに要素を追加した新しいクエリを返します。
// Returns a new query with an element appended to the query.
//
//	パラメータ:
//	- q: クエリ
//	- item: 追加する要素
//
//	Parameters:
//	- q: The query.
//	- item: The element to append.
//
func Append[T any](q Query[T], item T) Query[T] {
	return func(yield func(T) bool) {
		for v := range q {
			if !yield(v) {
				return
			}
		}
		yield(item)
	}
}

// Concat は、2 つのクエリを連結した新しいクエリを返します。
// Returns a new query that concatenates two queries.
//
//	パラメータ:
//	- q1: 連結する最初のクエリ
//	- q2: 連結する 2 番目のクエリ
//
//	Parameters:
//	- q1: The first query to concatenate.
//	- q2: The second query to concatenate.
//
func Concat[T any](q1, q2 Query[T]) Query[T] {
	return func(yield func(T) bool) {
		for v := range q1 {
			if !yield(v) {
				return
			}
		}
		for v := range q2 {
			if !yield(v) {
				return
			}
		}
	}
}

// Prepend は、クエリの先頭に要素を追加した新しいクエリを返します。
// Returns a new query with an element prepended to the query.
//
//	パラメータ:
//	- q: クエリ
//	- item: 先頭に追加する要素
//
//	Parameters:
//	- q: The query.
//	- item: The element to prepend.
//
func Prepend[T any](q Query[T], item T) Query[T] {
	return func(yield func(T) bool) {
		yield(item)
		for v := range q {
			if !yield(v) {
				return
			}
		}
	}
}
