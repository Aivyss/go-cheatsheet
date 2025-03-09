package collection

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/collection"
	"math"
	"testing"
)

func TestDistinctMap(t *testing.T) {
	// given
	type testStruct struct {
		v int
	}
	nums := []int{1, 2, 2, 4, 1, 6, 8, 5, 4, 3, 76, 78, 5, 3, 31}

	// when
	result := collection.DistinctMap(
		nums,
		func(num int) int {
			return num
		},
		func(num int) testStruct {
			return testStruct{v: num}
		},
	)

	// then
	expectedResult := []testStruct{
		{v: 1},
		{v: 2},
		{v: 4},
		{v: 6},
		{v: 8},
		{v: 5},
		{v: 3},
		{v: 76},
		{v: 78},
		{v: 31},
	}
	assert.Equal(t, expectedResult, result)
}

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

func TestFirstOrNil(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		// given
		type testStruct struct {
			flag   bool
			number int
		}
		slice := []testStruct{
			{
				flag:   false,
				number: 1,
			},
			{
				flag:   true,
				number: 2,
			},
			{
				flag:   true,
				number: 3,
			},
			{
				flag:   false,
				number: 4,
			},
		}

		// when
		v := collection.FirstOrNil(slice, func(t testStruct) bool {
			return t.flag
		})

		// then
		assert.NotNil(t, v)
		assert.Equal(t, 2, v.number)
	})
	t.Run("nil result", func(t *testing.T) {
		// given
		type testStruct struct {
			flag   bool
			number int
		}
		slice := []testStruct{
			{
				flag:   false,
				number: 1,
			},
			{
				flag:   false,
				number: 2,
			},
			{
				flag:   false,
				number: 3,
			},
			{
				flag:   false,
				number: 4,
			},
		}

		// when
		v := collection.FirstOrNil(slice, func(t testStruct) bool {
			return t.flag
		})

		// then
		assert.Nil(t, v)
	})
}

func TestLastOrNil(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		// given
		type testStruct struct {
			flag   bool
			number int
		}
		slice := []testStruct{
			{
				flag:   false,
				number: 1,
			},
			{
				flag:   true,
				number: 2,
			},
			{
				flag:   true,
				number: 3,
			},
			{
				flag:   false,
				number: 4,
			},
		}

		// when
		v := collection.LastOrNil(slice, func(t testStruct) bool {
			return t.flag
		})

		// then
		assert.NotNil(t, v)
		assert.Equal(t, 3, v.number)
	})

	t.Run("nil result", func(t *testing.T) {
		// given
		type testStruct struct {
			flag   bool
			number int
		}
		slice := []testStruct{
			{
				flag:   false,
				number: 1,
			},
			{
				flag:   false,
				number: 2,
			},
			{
				flag:   false,
				number: 3,
			},
			{
				flag:   false,
				number: 4,
			},
		}

		// when
		v := collection.LastOrNil(slice, func(t testStruct) bool {
			return t.flag
		})

		// then
		assert.Nil(t, v)
	})
}

func TestComplexSort(t *testing.T) {
	t.Run("trivial", func(t *testing.T) {
		// given
		type testStruct struct {
			order1 int
			order2 string
			order3 string
		}

		inputArr := []testStruct{
			{
				2,
				"a",
				"c",
			},
			{
				1,
				"a",
				"a",
			},
			{
				1,
				"b",
				"c",
			},
			{
				1,
				"b",
				"a",
			},
			{
				1,
				"a",
				"a",
			},
			{
				2,
				"a",
				"b",
			},
			{
				3,
				"b",
				"b",
			},
			{
				2,
				"a",
				"a",
			},
		}

		// when
		collection.ComplexSort(
			inputArr,
			func(i, j int) bool {
				return inputArr[i].order1 < inputArr[j].order1
			},
			func(i, j int) bool {
				return inputArr[i].order2 < inputArr[j].order2
			},
			func(i, j int) bool {
				return inputArr[i].order3 < inputArr[j].order3
			},
		)

		// then
		expectedResult := []testStruct{
			{1, "a", "a"},
			{1, "a", "a"},
			{1, "b", "a"},
			{1, "b", "c"},
			{2, "a", "a"},
			{2, "a", "b"},
			{2, "a", "c"},
			{3, "b", "b"},
		}
		assert.Equal(t, expectedResult, inputArr)
	})
	t.Run("no compare function", func(t *testing.T) {
		// given
		type testStruct struct {
			order1 int
			order2 string
			order3 string
		}

		inputArr := []testStruct{
			{
				2,
				"a",
				"c",
			},
			{
				1,
				"a",
				"a",
			},
			{
				1,
				"b",
				"c",
			},
			{
				1,
				"b",
				"a",
			},
			{
				1,
				"a",
				"a",
			},
			{
				2,
				"a",
				"b",
			},
			{
				3,
				"b",
				"b",
			},
			{
				2,
				"a",
				"a",
			},
		}

		// when
		collection.ComplexSort(inputArr)

		// then
		expectedResult := []testStruct{
			{
				2,
				"a",
				"c",
			},
			{
				1,
				"a",
				"a",
			},
			{
				1,
				"b",
				"c",
			},
			{
				1,
				"b",
				"a",
			},
			{
				1,
				"a",
				"a",
			},
			{
				2,
				"a",
				"b",
			},
			{
				3,
				"b",
				"b",
			},
			{
				2,
				"a",
				"a",
			},
		}
		assert.Equal(t, expectedResult, inputArr)
	})
	t.Run("with nil function", func(t *testing.T) {
		// given
		type testStruct struct {
			order1 int
			order2 string
			order3 string
		}

		inputArr := []testStruct{
			{
				2,
				"a",
				"c",
			},
			{
				1,
				"a",
				"a",
			},
			{
				1,
				"b",
				"c",
			},
			{
				1,
				"b",
				"a",
			},
			{
				1,
				"a",
				"a",
			},
			{
				2,
				"a",
				"b",
			},
			{
				3,
				"b",
				"b",
			},
			{
				2,
				"a",
				"a",
			},
		}

		// when
		collection.ComplexSort(
			inputArr,
			func(i, j int) bool {
				return inputArr[i].order1 < inputArr[j].order1
			},
			nil,
			func(i, j int) bool {
				return inputArr[i].order3 < inputArr[j].order3
			},
		)

		// then
		expectedResult := []testStruct{
			{1, "a", "a"},
			{1, "b", "a"},
			{1, "a", "a"},
			{1, "b", "c"},
			{2, "a", "a"},
			{2, "a", "b"},
			{2, "a", "c"},
			{3, "b", "b"},
		}
		assert.Equal(t, expectedResult, inputArr)
	})
}
