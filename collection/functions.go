package collection

import (
	"errors"
)

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
	return Distinct(arr, func(t T) T {
		return t
	})
}

func Distinct[T any, C comparable](arr []T, getComparable func(t T) C) []T {
	return DistinctMap(arr, getComparable, func(t T) T {
		return t
	})
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

func FirstOrNil[S ~[]T, T any](arr S, predicate func(t T) bool) *T {
	first, err := First(arr, predicate)
	if errors.Is(err, NoSliceElementErr) {
		return nil
	}

	return &first
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

func LastOrNil[S ~[]T, T any](arr S, predicate func(t T) bool) *T {
	last, err := Last(arr, predicate)
	if errors.Is(err, NoSliceElementErr) {
		return nil
	}

	return &last
}

func Map[A, B any](arr []A, convert func(a A) B) []B {
	convertedArr := make([]B, 0, len(arr))
	ForEach(arr, func(a A) {
		b := convert(a)
		convertedArr = append(convertedArr, b)
	})

	return convertedArr
}

func ListOf[T any](values ...T) []T {
	return values
}

func DistinctMap[T, V any, S ~[]T, C comparable](arr S, getComparable func(t T) C, converter func(t T) V) []V {
	isDuplicate := map[C]bool{}
	results := make([]V, 0, len(arr))
	for _, a := range arr {
		c := getComparable(a)
		if !isDuplicate[c] {
			isDuplicate[c] = true
			results = append(results, converter(a))
		}
	}

	return results
}
