package sort

import (
	"container/heap"

	"github.com/ryanrain2016/utils/collections"
)

type IHeap[T comparable] interface {
	ISort[T]
	Pop(index ...int) (T, error)
	Append(item T)
}

type Heap[T comparable] struct {
	less  []lessFunc[T]
	items IHeap[T]
}

func (h *Heap[T]) Copy() *Heap[T] {
	return &Heap[T]{
		less: h.less,
	}
}

func HeapBy[T comparable](less ...lessFunc[T]) *Heap[T] {
	if len(less) == 0 {
		panic("should has at least one less function")
	}
	return &Heap[T]{
		less:  less,
		items: collections.NewList[T](0),
	}
}

func (h *Heap[T]) Reverse() *Heap[T] {
	less := make([]lessFunc[T], len(h.less))
	for i, lf := range h.less {
		less[i] = func(t1, t2 T) bool {
			return !lf(t1, t2)
		}
	}
	return &Heap[T]{
		less:  less,
		items: collections.NewList[T](0),
	}
}

func (h *Heap[T]) Less(i, j int) bool {
	p, _ := h.items.GetItem(i)
	q, _ := h.items.GetItem(j)
	var k int
	for k = 0; k < len(h.less)-1; k++ {
		less := h.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return h.less[k](p, q)
}

func (h *Heap[T]) Len() int {
	return h.items.Len()
}

func (h *Heap[T]) Swap(i, j int) {
	h.items.Swap(i, j)
}

func (h *Heap[T]) Push(x any) {
	v := x.(T)
	h.items.Append(v)
}

func (h *Heap[T]) Pop() any {
	v, _ := h.items.Pop()
	return v
}

func (h *Heap[T]) Init(items IHeap[T]) {
	if items != nil {
		h.items = items
	}
	heap.Init(h)
}

func (h *Heap[T]) Fix(items IHeap[T], i int) {
	if items != nil {
		h.items = items
	}
	heap.Fix(h, i)
}

func (h *Heap[T]) HeapPop(items IHeap[T]) T {
	if items != nil {
		h.items = items
	}
	return heap.Pop(h).(T)
}

func (h *Heap[T]) HeapPush(items IHeap[T], t T) {
	if items != nil {
		h.items = items
	}
	heap.Push(h, t)
}

func (h *Heap[T]) Remove(items IHeap[T], i int) T {
	if items != nil {
		h.items = items
	}
	return heap.Remove(h, i).(T)
}

func (h *Heap[T]) HeapPushPop(items IHeap[T], t T) T {
	h.HeapPush(items, t)
	return h.HeapPop(items)
}

func (h *Heap[T]) Heapify(items IHeap[T]) {
	hh := h.Copy()
	hh.Init(items)
}

func (h *Heap[T]) HeapReplace(items IHeap[T], item T) T {
	r := h.HeapPop(items)
	h.HeapPush(items, item)
	return r
}

func (h *Heap[T]) NLargest(n int, items IHeap[T]) []T {
	r := make([]T, 0, n)
	hr := h.Reverse()
	if items != nil {
		l := collections.NewList[T](0)
		for i := 0; i < items.Len(); i++ {
			v, _ := items.GetItem(i)
			l.Append(v)
		}
		hr.Init(l)
	}
	for i := 0; i < n && i < items.Len(); i++ {
		r = append(r, hr.HeapPop(nil))
	}
	return r
}

func (h *Heap[T]) NSmallest(n int, items IHeap[T]) []T {
	r := make([]T, 0, n)
	hr := h.Copy()
	if items != nil {
		l := collections.NewList[T](0)
		for i := 0; i < items.Len(); i++ {
			v, _ := items.GetItem(i)
			l.Append(v)
		}
		hr.Init(l)
	}
	for i := 0; i < n && i < items.Len(); i++ {
		r = append(r, hr.HeapPop(nil))
	}
	return r
}
