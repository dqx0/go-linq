package main

import (
	"github.com/dqx0/linq"
)

func main() {
	list := []int{1, 4, 7, 2, 5, 8, 3, 6, 9}

	q := linq.New(list)
	r := linq.New([]int{})
	println(q.Any()) // true
	println(r.Any()) // false
	s := linq.Select(q, func(item int) int {
		return item * 2 // 2, 8, 14, 4, 10, 16, 6, 12, 18
	})
	s = linq.Select(s, func(item int) int {
		return item * 3 // 6, 24, 42, 12, 30, 48, 18, 36, 54
	})
	s = linq.Where(s, func(item int) bool {
		return item < 30 // 6, 24, 12, 18
	})
	s = linq.OrderBy(s, func(a, b int) int { // 6, 24, 12, 18
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	s = linq.OrderByDescending(s, func(a, b int) int { // 24, 18, 12, 6
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	u := s.ToSlice()
	for _, v := range u {
		println(v)
	}
	t := linq.Aggregate(s, 0, func(accumulator int, item int) int {
		return accumulator + item // 60
	})
	println(t) // 60
}
