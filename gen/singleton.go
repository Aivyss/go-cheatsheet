package gen

import (
	"go-cheatsheet/utils"
	"sync"
)

type SingletonSupplier[T any] interface {
	Get() T
}

type singletonSupplier[T any] struct {
	once     sync.Once
	t        *T
	supplier func() T
}

func (s *singletonSupplier[T]) Get() T {
	s.once.Do(func() {
		s.t = utils.Pointer(s.supplier())
	})

	return *s.t
}

func NewSingletonSupplier[T any](supplier func() T) SingletonSupplier[T] {
	return &singletonSupplier[T]{
		supplier: supplier,
	}
}
