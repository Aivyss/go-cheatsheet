package collection

import (
	"errors"
	"math"
)

type pair[T any] struct {
	idx int
	t   T
}

var NoSliceElementErr = errors.New("no_slice_element_error")

func ForEach[T any](arr []T, runnable func(t T)) {
	for _, t := range arr {
		runnable(t)
	}
}

func Filter[T any](arr []T, predicate func(t T) bool) []T {
	filteredArr := make([]T, 0, len(arr))
	for _, t := range arr {
		if predicate(t) {
			filteredArr = append(filteredArr, t)
		}
	}

	return filteredArr
}

func Count[T any](arr []T, predicate func(t T) bool) int {
	cnt := 0
	ForEach(arr, func(t T) {
		if predicate(t) {
			cnt += 1
		}
	})

	return cnt
}

func DistinctSimple[T comparable](arr []T) []T {
	set := map[T]bool{}

	return Filter(arr, func(t T) bool {
		flag, ok := set[t]
		if flag && ok {
			return false
		}

		set[t] = true
		return set[t]
	})
}

func Distinct[T any, C comparable](arr []T, getComparable func(t T) C) []T {
	collection := map[C][]pair[T]{}
	for i, t := range arr {
		c := getComparable(t)
		s := collection[c]
		if s == nil {
			s = []pair[T]{}
		}

		s = append(s, pair[T]{
			idx: i,
			t:   t,
		})
		collection[c] = s
	}

	filteredArr := make([]T, 0, len(arr))
	for {
		if len(collection) < 1 {
			break
		}

		var minC C
		var minT T
		var minI = math.MaxInt
		for c, pairs := range collection {
			p := pairs[0]
			if p.idx <= minI {
				minC = c
				minT = p.t
				minI = p.idx
			}
		}

		delete(collection, minC)
		filteredArr = append(filteredArr, minT)
	}

	return filteredArr
}

func First[T any](arr []T, predicate func(t T) bool) (T, error) {
	for _, t := range arr {
		if predicate(t) {
			return t, nil
		}
	}

	var tmp T
	return tmp, NoSliceElementErr
}

func Last[T any](arr []T, predicate func(t T) bool) (T, error) {
	for i := len(arr) - 1; i >= 0; i-- {
		if predicate(arr[i]) {
			return arr[i], nil
		}
	}

	var tmp T
	return tmp, NoSliceElementErr
}

func Map[A, B any](arr []A, convert func(a A) B) []B {
	convertedArr := make([]B, 0, len(arr))
	ForEach(arr, func(a A) {
		b := convert(a)
		convertedArr = append(convertedArr, b)
	})

	return convertedArr
}
