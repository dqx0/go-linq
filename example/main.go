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
	s := q.Select(func(item int) int {
		return item * 2 // 2, 8, 14, 4, 10, 16, 6, 12, 18
	}).Select(func(item int) int {
		return item * 3 // 6, 24, 42, 12, 30, 48, 18, 36, 54
	}).Where(func(item int) bool {
		return item < 30 // 6, 24, 12, 18
	}).Aggregate(0, func(accumulator int, item int) int {
		return accumulator + item // 60
	})
	println(s) // 60
}
