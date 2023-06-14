package utils

import "sync"

// FIXME add test code

type ConcurrentSlice[T any] struct {
	sync.RWMutex
	items []T
}

type ConcurrentSliceItem[T any] struct {
	Index int
	Value T
}

func (cs *ConcurrentSlice[T]) Append(item T) {
	cs.Lock()
	defer cs.Unlock()

	cs.items = append(cs.items, item)
}

// Iter Iterates over the items in the concurrent slice
// Each item is sent over a channel, so that
// we can iterate over the slice using the builin range keyword
func (cs *ConcurrentSlice[T]) Iter() <-chan ConcurrentSliceItem[T] {
	c := make(chan ConcurrentSliceItem[T])

	f := func() {
		cs.Lock()
		defer cs.Unlock()
		for index, value := range cs.items {
			c <- ConcurrentSliceItem[T]{index, value}
		}
		close(c)
	}
	go f()

	return c
}

func (cs *ConcurrentSlice[T]) Length() int {
	return len(cs.items)
}

func (cs *ConcurrentSlice[T]) GetItem() []T {
	return cs.items
}

func Contains[T any](arr []T, equal func(value T) bool) bool {
	for _, item := range arr {
		if equal(item) {
			return true
		}
	}
	return false
}

// MapValue 순서보장 O, 중복 제거 X
func MapValue[T any, V any](slice []T, getValue func(element T) V) []V {
	valueList := make([]V, 0, len(slice))

	for _, element := range slice {
		valueList = append(valueList, getValue(element))
	}

	return valueList
}

// MapValueUnique 중복 제거 O, 순서 보장 O
func MapValueUnique[T any, V comparable](slice []T, getValue func(element T) V) []V {
	valueList := make([]V, 0, len(slice))
	checker := make(map[V]bool)

	for _, element := range slice {
		value := getValue(element)
		if !checker[value] {
			valueList = append(valueList, value)
			checker[value] = true
		}
	}

	return valueList
}
