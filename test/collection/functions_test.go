package collection

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/collection"
	"math"
	"testing"
)

func TestDistinct(t *testing.T) {
	type nonComparable struct {
		num   int
		dummy func()
	}

	arr := []nonComparable{
		{
			num: 1,
		},
		{
			num: 3,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 3,
		},
	}
	expectedResult := []nonComparable{
		{
			num: 1,
		},
		{
			num: 3,
		},
		{
			num: 2,
		},
	}
	actualResult := collection.Distinct(arr, func(t nonComparable) int {
		return t.num
	})
	t.Log("actualResult=", actualResult)

	assert.Equal(t, 3, len(actualResult))
	for i, actual := range actualResult {
		assert.Equal(t, expectedResult[i].num, actual.num)
	}
}

func TestDistinctSimple(t *testing.T) {
	arr := []int{
		1,
		3,
		2,
		2,
		2,
		1,
		1,
		3,
	}
	expectedResult := []int{
		1,
		3,
		2,
	}
	actualResult := collection.DistinctSimple(arr)
	t.Log("actualResult=", actualResult)

	assert.Equal(t, 3, len(actualResult))
	for i, actual := range actualResult {
		assert.Equal(t, expectedResult[i], actual)
	}
}

func TestFilter(t *testing.T) {
	// given
	type nonComparable struct {
		num   int
		dummy func()
	}

	arr := []nonComparable{
		{
			num: 1,
		},
		{
			num: 3,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 3,
		},
	}

	// when1
	actuals := collection.Filter(arr, func(t nonComparable) bool {
		return t.num == 1
	})

	// then1
	assert.Equal(t, 4, len(actuals))

	// when2
	actuals2 := collection.Filter(arr, func(t nonComparable) bool {
		return t.num == 5
	})

	// then2
	assert.Equal(t, 0, len(actuals2))
}

func TestCount(t *testing.T) {
	// given
	type nonComparable struct {
		num   int
		dummy func()
	}

	arr := []nonComparable{
		{
			num: 1,
		},
		{
			num: 3,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 3,
		},
	}

	// when1
	actuals := collection.Count(arr, func(t nonComparable) bool {
		return t.num == 1
	})

	// then1
	assert.Equal(t, 4, actuals)

	// when2
	actuals2 := collection.Count(arr, func(t nonComparable) bool {
		return t.num == 5
	})

	// then2
	assert.Equal(t, 0, actuals2)
}

func TestFirst(t *testing.T) {
	// given
	type nonComparable struct {
		num   int
		dummy func()
	}

	arr := []nonComparable{
		{
			num: 1,
		},
		{
			num: 3,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 3,
		},
	}

	// when1
	actual, err := collection.First(arr, func(t nonComparable) bool {
		return t.num == 1
	})

	// then1
	assert.Nil(t, err)
	assert.Equal(t, 1, actual.num)

	// when2
	_, err = collection.First(arr, func(t nonComparable) bool {
		return t.num == 5
	})

	// then2
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, collection.NoSliceElementErr))
}

func TestLast(t *testing.T) {
	// given
	type nonComparable struct {
		num   int
		dummy func()
	}

	arr := []nonComparable{
		{
			num: 1,
		},
		{
			num: 3,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 3,
		},
	}

	// when1
	actual, err := collection.Last(arr, func(t nonComparable) bool {
		return t.num == 1
	})

	// then1
	assert.Nil(t, err)
	assert.Equal(t, 1, actual.num)

	// when2
	_, err = collection.Last(arr, func(t nonComparable) bool {
		return t.num == 5
	})

	// then2
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, collection.NoSliceElementErr))
}

func TestMap(t *testing.T) {
	// given
	type nonComparable struct {
		num   int
		dummy func()
	}

	arr := []nonComparable{
		{
			num: 1,
		},
		{
			num: 3,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 2,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 1,
		},
		{
			num: 3,
		},
	}
	expected := []int{
		1,
		3,
		2,
		2,
		2,
		1,
		1,
		1,
		3,
	}

	// when
	actuals := collection.Map(arr, func(t nonComparable) int {
		return t.num
	})

	// then
	assert.Equal(t, expected, actuals)
}

func TestListOf(t *testing.T) {
	t.Run("any", func(t *testing.T) {
		// given
		value1 := 1
		value2 := "str"
		value3 := true
		value4 := math.Pi
		value5 := float32(3.14)
		value6 := int32(2)
		value7 := int64(2)
		expected := []any{value1, value2, value3, value4, value5, value6, value7}

		// when
		actual := collection.ListOf[any](value1, value2, value3, value4, value5, value6, value7)

		// then
		assert.Equal(t, expected, actual)
	})
	t.Run("typed", func(t *testing.T) {
		type testStruct struct {
			value string
		}

		v1 := testStruct{value: "a"}
		v2 := testStruct{value: "b"}
		v3 := testStruct{value: "c"}
		v4 := testStruct{value: "d"}
		v5 := testStruct{value: "e"}
		expected := []testStruct{v1, v2, v3, v4, v5}

		// when
		actual := collection.ListOf(v1, v2, v3, v4, v5)

		// then
		assert.Equal(t, expected, actual)
	})
}
