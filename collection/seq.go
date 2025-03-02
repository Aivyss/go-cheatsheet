package collection

import (
	"iter"
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

func LoadSeq[T any](s iter.Seq[T]) []T {
	var result []T
	for a := range s {
		result = append(result, a)
	}

	return result
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

func FilterSeq[T any](s iter.Seq[T], predicate func(t T) bool) iter.Seq[T] {
	return func(yield func(a T) bool) {
		for a := range s {
			if predicate(a) {
				if !yield(a) {
					return
				}
			}
		}
	}
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
