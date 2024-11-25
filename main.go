package main

import (
	"fmt"
)

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

// func (s Stack[T]) Contains(val T) bool {
// 	for _, v := range s.vals {
// 		if v == val {
// 			return true
// 		}
// 	}
// 	return false
// }

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func main() {

	// Wrong Logic,
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	// intStack.Pop()
	// intStack.Pop()
	v, ok := intStack.Pop()
	fmt.Println(v, ok)

	numbers := []int{1, 2, 3, 4, 5}
	sum := Reduce(numbers, 0, func(acc, num int) int {
		return acc + num
	})
	fmt.Println(sum) // Output: 15
	// Generics
	result := MapValues([]int{1, 2, 3, 4}, func(n int) int {
		return n * 3
	})
}
