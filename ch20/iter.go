package ch20

import "fmt"

type Iterator[T any] interface {
	Next() bool // 是否存在下一个
	First() T
	CurrentItem() T
}

type BusIterator[T any] struct {
	elements []T
	index    int
	count    int
}

func NewBusIterator[T any](data []T) *BusIterator[T] {
	return &BusIterator[T]{
		elements: data,
		index:    -1,
		count:    len(data),
	}
}

func (i *BusIterator[T]) Next() bool {
	if i.index == i.count-1 {
		return false
	}
	i.index++
	return true
}

func (i *BusIterator[T]) First() T {
	var d T
	if i.count <= 0 {
		return d
	}
	return i.elements[0]
}

func (i *BusIterator[T]) CurrentItem() T {
	var d T
	if i.count <= i.index {
		return d
	}
	return i.elements[i.index]
}

func IterMain() {
	bus := NewBusIterator[int]([]int{1, 2, 3, 4, 5, 6})
	for bus.Next() {
		fmt.Println("item:", bus.CurrentItem())
	}
}
