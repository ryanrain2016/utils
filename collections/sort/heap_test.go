package sort

import (
	"container/heap"
	"reflect"
	"testing"

	"github.com/ryanrain2016/utils/collections"
	"github.com/stretchr/testify/assert"
)

func TestHeapInit(t *testing.T) {
	var h = HeapBy(func(i1, i2 int) bool {
		return i1 < i2
	})
	l := collections.NewList(0, 0, 3, 6, 2, 4, 5, 1, 7, 8, 9)
	h.Heapify(l)
	sl := l.ToSlice()
	expected := []int{0, 2, 1, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(sl, expected) {
		t.Errorf("Init Failed. expect %v, but got %v", expected, sl)
	}
}

func TestPushAndPop(t *testing.T) {
	h := HeapBy(func(x, y int) bool { return x < y })
	h.Push(3)
	h.Push(1)
	h.Push(4)
	v1 := h.Pop().(int)
	v2 := h.Pop().(int)
	v3 := h.Pop().(int)
	if v1 != 4 || v2 != 1 || v3 != 3 {
		t.Error("TestPushAndPop failed")
	}
}

func TestHeapPushAndHeapPop(t *testing.T) {
	h := HeapBy(func(x, y int) bool { return x < y })
	items_ := h.items
	items := []int{3, 1, 4}
	for _, item := range items {
		h.HeapPush(items_, item)
	}
	v1 := h.HeapPop(items_)
	v2 := h.HeapPop(items_)
	v3 := h.HeapPop(items_)
	if v1 != 1 || v2 != 3 || v3 != 4 {
		t.Error("TestHeapPushAndHeapPop failed")
	}
}

func TestHeapPushPop(t *testing.T) {
	h := HeapBy(func(x, y int) bool { return x < y })
	items := collections.NewList(0, 3, 1, 4)
	h.Heapify(items)
	v := h.HeapPushPop(items, 2)
	if v != 1 {
		t.Errorf("TestHeapPushPop failed, expect 1, but got %v", v)
	}
	sl := items.ToSlice()
	if !reflect.DeepEqual(sl, []int{2, 3, 4}) {
		t.Errorf("TestHeapPushPop failed, expect [2 3 4], but got %v", sl)
	}
}

func TestFix(t *testing.T) {
	h := HeapBy(func(x, y int) bool { return x < y })
	items := collections.NewList(0, 3, 1, 4)
	h.Heapify(items)
	items.SetItem(1, 6)
	h.Fix(items, 1)
	v1 := heap.Pop(h).(int)
	v2 := heap.Pop(h).(int)
	v3 := heap.Pop(h).(int)
	if v1 != 1 || v2 != 4 || v3 != 6 {
		t.Errorf("TestFix failed, expected 1, 4, 6, but got [%v %v %v]", v1, v2, v3)
	}
}

func TestHeapify(t *testing.T) {
	h := HeapBy(func(x, y int) bool { return x < y })
	items := collections.NewList(0, 3, 1, 4)
	h.Heapify(items)
	v1, _ := items.Pop(0)
	v2, _ := items.Pop(0)
	v3, _ := items.Pop(0)
	if v1 != 1 || v2 != 3 || v3 != 4 {
		t.Error("TestHeapify failed")
	}
}

func TestHeapReplace(t *testing.T) {
	h := HeapBy(func(x, y int) bool { return x < y })
	items := collections.NewList(0, 3, 1, 4)
	h.Heapify(items)
	v := h.HeapReplace(items, 2)
	if v != 1 {
		t.Error("TestHeapReplace failed")
	}
}

func TestPopEmpty(t *testing.T) {
	assert.Panics(t, func() {
		h := HeapBy(func(x, y int) bool { return x < y })
		h.HeapPop(collections.NewList[int](0))
	})
}

func TestHeapNLargest(t *testing.T) {
	h := HeapBy(func(x, y int) bool { return x < y })
	items := collections.NewList(0, 3, 1, 4, 5, 2, 9, 7, 8, 6)
	sl := h.NLargest(4, items)
	if !reflect.DeepEqual(sl, []int{9, 8, 7, 6}) {
		t.Errorf("HeapNLargest error. expect [9 8 7 6], but got %v", sl)
	}
}

func TestHeapNSmallest(t *testing.T) {
	h := HeapBy(func(x, y int) bool { return x < y })
	items := collections.NewList(0, 3, 1, 4, 5, 2, 9, 7, 8, 6)
	sl := h.NSmallest(4, items)
	if !reflect.DeepEqual(sl, []int{1, 2, 3, 4}) {
		t.Errorf("HeapNLargest error. expect [0 1 2 3], but got %v", sl)
	}
}
