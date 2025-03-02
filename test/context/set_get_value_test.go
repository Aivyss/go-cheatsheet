package context

import (
	"context"
	"github.com/stretchr/testify/assert"
	util_ctx "go-cheatsheet/context"
	"testing"
)

func TestSetByCtxGetByCtx(t *testing.T) {
	// prepared
	type testStruct struct {
		value string
	}

	t.Run("get with no value", func(t *testing.T) {
		// given
		var testKey util_ctx.ContextKey[testStruct] = "test-key"
		ctx := context.Background()

		// when
		value, err := util_ctx.GetByCtx(ctx, testKey)

		// then
		assert.Equal(t, util_ctx.ErrNoContextRecord, err)
		assert.Equal(t, testStruct{}, value)
	})
	t.Run("get with a value", func(t *testing.T) {
		// given
		var testKey util_ctx.ContextKey[testStruct] = "test-key"
		ctx := context.Background()
		expectedValue := testStruct{
			value: "test-struct",
		}

		// when 1
		wrappedCtx := util_ctx.SetByCtx(ctx, testKey, &expectedValue)
		// then 1
		assert.NotEqual(t, ctx, wrappedCtx)

		// when 2
		actualValue, err := util_ctx.GetByCtx(wrappedCtx, testKey)
		// then 2
		assert.Nil(t, err)
		assert.Equal(t, expectedValue, actualValue)
	})
	t.Run("set with nil value", func(t *testing.T) {
		// given
		var testKey util_ctx.ContextKey[testStruct] = "test-key"
		ctx := context.Background()
		var value *testStruct

		// when 1
		wrappedCtx := util_ctx.SetByCtx(ctx, testKey, value)
		// then 1
		assert.Equal(t, ctx, wrappedCtx)

		// when 2
		_, err := util_ctx.GetByCtx(wrappedCtx, testKey)
		// then 2
		assert.NotNil(t, err)
		assert.Equal(t, util_ctx.ErrNoContextRecord, err)
	})
}
