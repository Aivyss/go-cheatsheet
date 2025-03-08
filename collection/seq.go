package collection

import (
	"iter"
	"math"
)

func SeqOf[T any](values ...T) iter.Seq[T] {
	return func(yield func(a T) bool) {
		for _, a := range values {
			if !yield(a) {
				return
			}
		}
	}
}

func LoadSeqN[T any](s iter.Seq[T], n int) []T {
	var result []T
	count := 0
	if n < 0 {
		n = math.MaxInt
	}

	for a := range s {
		result = append(result, a)
		count += 1
		if count == n {
			break
		}
	}

	return result
}

func LoadSeq[T any](s iter.Seq[T]) []T {
	return LoadSeqN(s, -1)
}

func MapSeq[A, B any](s iter.Seq[A], convert func(a A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for a := range s {
			if !yield(convert(a)) {
				return
			}
		}
	}
}

func FilterSeqN[T any](s iter.Seq[T], predicate func(t T) bool, n int) iter.Seq[T] {
	return func(yield func(a T) bool) {
		count := 0
		if n < 0 {
			n = math.MaxInt
		}

		for a := range s {
			if predicate(a) {
				if !yield(a) {
					return
				}

				count += 1
				if count == n {
					return
				}
			}
		}
	}
}

func FilterSeq[T any](s iter.Seq[T], predicate func(t T) bool) iter.Seq[T] {
	return FilterSeqN(s, predicate, -1)
}

func CountSeq[T any](s iter.Seq[T]) int {
	var count int
	for range s {
		count++
	}

	return count
}

func ConcatSeq[T any](s iter.Seq[T], values ...T) iter.Seq[T] {
	return func(yield func(a T) bool) {
		for a := range s {
			if !yield(a) {
				return
			}
		}
		for _, a := range values {
			if !yield(a) {
				return
			}
		}
	}
}
