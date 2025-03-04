package collection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/collection"
	"testing"
)

func TestMapSeq(t *testing.T) {
	// given
	arr := []int{0, 1, 2, 3}
	original := collection.SeqOf(arr...)

	// when
	strSeq := collection.MapSeq(original, func(i int) string {
		return fmt.Sprint(i)
	})

	// then 1
	var thenResult1 []string
	for s := range strSeq {
		thenResult1 = append(thenResult1, s)
		if s == "2" {
			break
		}
	}
	assert.Equal(t, []string{"0", "1", "2"}, thenResult1)

	// then 2
	var thenResult2 []string
	for s := range strSeq {
		thenResult2 = append(thenResult2, s)
	}
	assert.Equal(t, []string{"0", "1", "2", "3"}, thenResult2)
}

func TestSeqOf(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		// when
		seq := collection.SeqOf[string]()

		// then
		for range seq {
			t.Fail()
		}
	})
	t.Run("not empty", func(t *testing.T) {
		// given
		expectedResult := []string{"1", "2", "3"}

		// when
		seq := collection.SeqOf(1, 2, 3)

		// then
		actualResult := make([]string, 0, len(expectedResult))
		for s := range seq {
			actualResult = append(actualResult, fmt.Sprint(s))
		}
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestConcatSeq(t *testing.T) {
	t.Run("no concat", func(t *testing.T) {
		// given
		seq := collection.SeqOf(1, 2, 3)

		// when
		newSeq := collection.ConcatSeq(seq)

		// then
		actualResult := make([]int, 0, 3)
		for s := range newSeq {
			actualResult = append(actualResult, s)
		}
		assert.Equal(t, []int{1, 2, 3}, actualResult)
	})

	t.Run("concat one", func(t *testing.T) {
		// given
		seq := collection.SeqOf(1, 2, 3)

		// when
		newSeq := collection.ConcatSeq(seq, 4)

		// then
		actualResult := make([]int, 0, 4)
		for s := range newSeq {
			actualResult = append(actualResult, s)
		}
		assert.Equal(t, []int{1, 2, 3, 4}, actualResult)
	})
	t.Run("concat many - 1", func(t *testing.T) {
		// given
		seq := collection.SeqOf(1, 2, 3)

		// when
		newSeq := collection.ConcatSeq(seq, 4, 5, 6)

		// then
		actualResult := make([]int, 0, 6)
		for s := range newSeq {
			actualResult = append(actualResult, s)
		}
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, actualResult)
	})
	t.Run("concat many - 2", func(t *testing.T) {
		// given
		seq := collection.SeqOf(1, 2, 3)

		// when
		newSeq1 := collection.ConcatSeq(seq, 4)
		newSeq2 := collection.ConcatSeq(newSeq1, 5)
		newSeq3 := collection.ConcatSeq(newSeq2, 6)

		// then
		actualResult := make([]int, 0, 6)
		for s := range newSeq3 {
			actualResult = append(actualResult, s)
		}
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, actualResult)
	})
}

func TestLoadSeq(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		// given
		seq := collection.SeqOf[int]()

		// when
		result := collection.LoadSeq(seq, -1)

		// then
		assert.Empty(t, result)
	})
	t.Run("not empty", func(t *testing.T) {
		// given
		seq := collection.SeqOf(1, 2, 3)

		// when
		result := collection.LoadSeq(seq, -1)

		// then
		assert.Equal(t, []int{1, 2, 3}, result)
	})
	t.Run("limit", func(t *testing.T) {
		// given
		seq := collection.SeqOf(1, 2, 3)

		// when
		result := collection.LoadSeq(seq, 2)

		// then
		assert.Equal(t, []int{1, 2}, result)
	})
}

func TestFilterSeq(t *testing.T) {
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
	seq := collection.SeqOf(arr...)

	// when1
	actuals := collection.FilterSeq(seq, func(t nonComparable) bool {
		return t.num == 1
	})

	// then1
	assert.Equal(t, 4, collection.CountSeq(actuals))

	// when2
	actuals2 := collection.FilterSeq(seq, func(t nonComparable) bool {
		return t.num == 5
	})

	// then2
	assert.Equal(t, 0, collection.CountSeq(actuals2))
}

func TestFilterN(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		// given
		seq := collection.SeqOf(1, 1, 1)

		// when
		actuals := collection.FilterSeqN(seq, func(t int) bool {
			return t == 1
		}, 2)

		// then
		assert.Equal(t, 2, collection.CountSeq(actuals))
	})
	t.Run("first N", func(t *testing.T) {
		// given
		seq := collection.SeqOf(1, 9102, 12934)

		// when
		actuals := collection.FilterSeqN(seq, func(t int) bool {
			return true
		}, 2)

		// then
		assert.Equal(t, 2, collection.CountSeq(actuals))
		assert.Equal(t, []int{1, 9102}, collection.LoadSeq(actuals, -1))
	})
}
