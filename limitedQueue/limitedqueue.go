package limitedqueue

import (
	"slices"
)

type LimitedQueue[T any] struct {
	capacity       int
	items          []T
	comparisonFunc func(T, T) int
}

func CreateLimitedQueue[T any](capacity int, comparisonFunc func(T, T) int) LimitedQueue[T] {
	return LimitedQueue[T]{
		capacity:       capacity,
		items:          make([]T, 0, capacity),
		comparisonFunc: comparisonFunc,
	}
}

func (q *LimitedQueue[T]) GetItems() []T {
	return q.items
}

func (q *LimitedQueue[T]) Insert(item T) {
	index, _ := slices.BinarySearchFunc(q.items, item, q.comparisonFunc)
	// if we have room to insert then do so.
	// If we don't have room to insert and would be inserting at the end of the slice, then do nothing
	// If we don't have room to insert and would be inserting anywhere other than the end of the slice
	// insert at desired position and remove last element.
	if len(q.items) == q.capacity {
		//we have no room to nsert
		if index != q.capacity-1 {
			q.items = slices.Insert(q.items, index, item)
			q.items = q.items[:q.capacity]
		}
	} else {
		q.items = slices.Insert(q.items, index, item)
	}
}
