package linq

// Contains はクエリに指定された要素が含まれているかどうかを返します。
// 含まれている場合は true を返し、含まれていない場合は false を返します。
//
// Contains returns true if the specified element is contained in the query,
// otherwise returns false.
//
// 	パラメータ:
// 	- q: クエリ
// 	- element: 検索する要素
//
// 	戻り値:
// 	- クエリに要素が含まれている場合は true、含まれていない場合は false
//
// 	Parameters:
// 	- q: The query.
// 	- element: The element to search for.
//
// 	Returns:
// 	- true if the element is contained in the query, otherwise false.
//
func Contains[T comparable](q Query[T], element T) bool {
	for v := range q {
		if v == element {
			return true
		}
	}
	return false
}
