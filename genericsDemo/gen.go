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

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

/*
interfaces and generics
*/

type CustomInteger interface {
	int | int16 | int32 | int64 | int8 | float32 | float64 | uint
}

/*
Map Values in Generics...
*/

func MapValues[T ~int | float64](values []T, mapfunc func(T) T) []T {
	var resultValues []T
	for _, v := range values {
		newVal := mapfunc(v)
		resultValues = append(resultValues, newVal)
	}
	return resultValues
}

func main() {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	v, ok := intStack.Pop()
	fmt.Println(v, ok)
	result := MapValues([]float64{1.1, 2.3, 3.3, 4.3}, func(n float64) float64 {
		return n * 3
	})
	fmt.Println(result)
}
