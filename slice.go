package utils

import "sync"

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
