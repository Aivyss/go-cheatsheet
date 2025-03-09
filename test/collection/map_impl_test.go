package collection

import (
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/collection"
	"go-cheatsheet/types"
	"slices"
	"sort"
	"testing"
)

func TestMultiValueMap(t *testing.T) {
	t.Run("no data1", func(t *testing.T) {
		// given
		m := collection.NewMultiValueMap[int, string]()

		// when
		values, ok := m.Get(1)

		// then
		assert.Empty(t, values)
		assert.True(t, m.IsEmpty())
		assert.False(t, ok)
		assert.Empty(t, 0, m.Len())
	})

	t.Run("Put", func(t *testing.T) {
		// given
		m := collection.NewMultiValueMap[int, string]()

		// when
		m.Put(1, "a")
		m.Put(1, "b")
		m.Put(1, "c")

		// then
		values, _ := m.Get(1)
		assert.Equal(t, []string{"a", "b", "c"}, values)
	})

	t.Run("Sort", func(t *testing.T) {
		// given
		m := collection.NewMultiValueMap[int, string]()
		inputValues := []string{"b", "a", "c"}
		orderBy := func(values []string, i, j int) bool {
			return values[i] < values[j]
		}

		// when
		for _, value := range inputValues {
			m.Put(1, value)
			m.Put(2, value)
		}

		// then
		m.Sort(1, orderBy)
		key1Values, _ := m.Get(1)
		key2Values, _ := m.Get(2)
		assert.Equal(t, []string{"a", "b", "c"}, key1Values)
		assert.Equal(t, inputValues, key2Values)
	})

	t.Run("SortAll", func(t *testing.T) {
		// given
		m := collection.NewMultiValueMap[int, string]()
		inputValues := []string{"b", "a", "c"}
		orderBy := func(values []string, i, j int) bool {
			return values[i] < values[j]
		}

		// when
		for _, value := range inputValues {
			m.Put(1, value)
			m.Put(2, value)
		}

		// then
		m.SortAll(orderBy)
		key1Values, _ := m.Get(1)
		key2Values, _ := m.Get(2)
		assert.Equal(t, []string{"a", "b", "c"}, key1Values)
		assert.Equal(t, []string{"a", "b", "c"}, key2Values)
	})

	t.Run("Keys", func(t *testing.T) {
		// given
		m := collection.NewMultiValueMap[int, string]()
		inputValues := []string{"b", "a", "c"}

		// when
		for _, value := range inputValues {
			m.Put(1, value)
			m.Put(2, value)
		}

		// then
		keys := collection.LoadSeq(m.Keys())
		slices.Sort(keys)
		assert.Equal(t, []int{1, 2}, keys)
	})
	t.Run("Entires", func(t *testing.T) {
		// given
		m := collection.NewMultiValueMap[int, string]()
		inputValues := []string{"b", "a", "c"}

		// when
		for _, value := range inputValues {
			m.Put(1, value)
			m.Put(2, value)
		}

		// then
		s := collection.LoadSeq2(m.Entries())
		sort.Slice(s, func(i, j int) bool {
			return s[i].First() < s[j].First()
		})
		assert.Equal(t, []types.Pair[int, []string]{
			types.PairOf(1, inputValues),
			types.PairOf(2, inputValues),
		}, s)
	})
}

func TestMultiValueMapForCoverage(t *testing.T) {
	t.Run("multiValueMapImpl.Sort quick stop", func(t *testing.T) {
		m := collection.NewMultiValueMap[int, string]()
		m.Sort(1, func(values []string, i, j int) bool {
			return values[i] < values[j]
		})
	})
}

func TestSet(t *testing.T) {
	t.Run("PutAll / Keys", func(t *testing.T) {
		// given
		expected := []int{1, 2, 3}
		s := collection.NewSet[int]()

		// when
		s.PutAll(1, 2, 3)

		// then
		assert.Equal(t, 3, s.Len())
		keys := collection.LoadSeq(s.Keys())
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
		assert.Equal(t, expected, keys)
	})
	t.Run("Put / Keys", func(t *testing.T) {
		// given
		expected := []int{1, 2, 3}
		s := collection.NewSet[int]()

		// when
		for _, v := range expected {
			s.Put(v)
		}

		// then
		assert.Equal(t, 3, s.Len())
		keys := collection.LoadSeq(s.Keys())
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
		assert.Equal(t, expected, keys)
	})
	t.Run("Contains", func(t *testing.T) {
		// given
		s := collection.NewSet[int]()
		s.PutAll(1, 2, 3)

		// when then 1
		assert.True(t, s.Contains(1))
		// when then 2
		assert.True(t, s.Contains(2))
		// when then 3
		assert.True(t, s.Contains(3))
		// when then 4
		assert.False(t, s.Contains(4))
	})
	t.Run("Len / IsEmpty", func(t *testing.T) {
		// given
		s := collection.NewSet[int]()

		// when then 1
		assert.True(t, s.IsEmpty())
		assert.Equal(t, 0, s.Len())
		// when
		s.PutAll(1, 2, 3)
		// when then 2
		assert.False(t, s.IsEmpty())
		assert.Equal(t, 3, s.Len())
	})
}
