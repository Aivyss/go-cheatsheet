package collection

import (
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/collection"
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
		assert.False(t, ok)
		assert.Empty(t, 0, m.Len())
	})

	t.Run("no data2", func(t *testing.T) {
		// given
		m := collection.NewMultiValueMap[int, string]()

		// when
		m.Put(1)

		// then
		values, ok := m.Get(1)
		assert.Empty(t, values)
		assert.True(t, ok)
	})

	t.Run("Put", func(t *testing.T) {
		// given
		m := collection.NewMultiValueMap[int, string]()

		// when
		m.Put(1, "a", "b", "c")

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
		m.Put(1, inputValues...)
		m.Put(2, inputValues...)

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
		m.Put(1, inputValues...)
		m.Put(2, inputValues...)

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
		m.Put(1, inputValues...)
		m.Put(2, inputValues...)
		m.Put(3)

		// then
		keys := m.Keys()
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
		assert.Equal(t, []int{1, 2, 3}, keys)
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
