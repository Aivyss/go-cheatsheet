package collection

import "sort"

type MultiValueMap[K comparable, V any] interface {
	Put(key K, values ...V)
	Get(key K) ([]V, bool)
	Sort(key K, orderBy func(values []V, i, j int) bool)
	SortAll(orderBy func(values []V, i, j int) bool)
	Keys() []K
	Len() int
}

type multiValueMapImpl[K comparable, V any] map[K][]V

func NewMultiValueMap[K comparable, V any]() MultiValueMap[K, V] {
	return multiValueMapImpl[K, V](map[K][]V{})
}

func (m multiValueMapImpl[K, V]) Put(key K, values ...V) {
	list, ok := m[key]
	if !ok {
		list = make([]V, 0, len(values))
	}

	list = append(list, values...)

	m[key] = list
}

func (m multiValueMapImpl[K, V]) Get(key K) ([]V, bool) {
	values, ok := m[key]
	return values, ok
}

func (m multiValueMapImpl[K, V]) Sort(key K, orderBy func(values []V, i int, j int) bool) {
	list, ok := m.Get(key)
	if !ok {
		return
	}

	sort.Slice(list, func(i, j int) bool {
		return orderBy(list, i, j)
	})
	m[key] = list
}

func (m multiValueMapImpl[K, V]) SortAll(orderBy func(values []V, i int, j int) bool) {
	for key := range m {
		m.Sort(key, orderBy)
	}
}

func (m multiValueMapImpl[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func (m multiValueMapImpl[K, V]) Len() int {
	return len(m)
}
