package context

import (
	"context"
	"errors"
)

type ContextKey[T any] string

var ErrNoContextRecord = errors.New("no context record")

func SetByCtx[T any](ctx context.Context, key ContextKey[T], value *T) context.Context {
	if value == nil {
		return ctx
	}

	return context.WithValue(ctx, key, value)
}

func GetByCtx[T any](ctx context.Context, key ContextKey[T]) (T, error) {
	var dummyT T
	obj := ctx.Value(key)
	if obj == nil {
		return dummyT, ErrNoContextRecord
	}

	t, ok := obj.(*T)
	if !ok {
		return dummyT, ErrNoContextRecord
	}
	if t == nil {
		return dummyT, ErrNoContextRecord
	}

	return *t, nil
}
