package collections

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListAppend(t *testing.T) {
	l := LinkedList[int]{}
	l.Append(1)
	if l.Len() != 1 {
		t.Errorf("Expected length of 1, but got %d", l.Len())
	}
	v, err := l.GetItem(0)
	assert.Nil(t, err)
	if v != 1 {
		t.Errorf("Expected item at index 0 to be 1, but got %v", v)
	}
}

func TestLinkedListClear(t *testing.T) {
	l := NewLinkedList(0, 1, 2, 3)
	l.Clear()
	if l.Len() != 0 {
		t.Errorf("Expected length of 0, but got %d", l.Len())
	}
}

func TestLinkedListCopy(t *testing.T) {
	l1 := NewLinkedList(0, 1, 2, 3)
	l2 := l1.Copy().(*LinkedList[int])
	if l1.Len() != l2.Len() {
		t.Errorf("Expected length of %d, but got %d", l1.Len(), l2.Len())
	}
	sl1 := l1.ToSlice()
	sl2 := l2.ToSlice()
	for i := 0; i < l1.Len(); i++ {
		if sl1[i] != sl2[i] {
			t.Errorf("Expected item at index %d to be %v, but got %v", i, sl1[i], sl2[i])
		}
	}
}

func TestLinkedListCount(t *testing.T) {
	l := NewLinkedList(0, 1, 2, 3, 2, 2)
	if l.Count(2) != 3 {
		t.Errorf("Expected count of 3, but got %d", l.Count(2))
	}
}

func TestLinkedListExtend(t *testing.T) {
	l1 := NewLinkedList(1, 2, 3)
	l2 := NewLinkedList(4, 5, 6)
	l1.Extend(l2)
	if l1.Len() != 6 {
		t.Errorf("Expected length of 6, but got %d", l1.Len())
	}
	sl1 := l1.ToSlice()
	sl2 := l2.ToSlice()
	for i := 0; i < l2.Len(); i++ {
		if sl1[i+3] != sl2[i] {
			t.Errorf("Expected item at index %d to be %v, but got %v", i+3, sl2[i], sl1[i+3])
		}
	}
}

func TestLinkedListExtendSlice(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	s := []int{4, 5, 6}
	l.Extend(s)
	if l.Len() != 6 {
		t.Errorf("Expected length of 6, but got %d", l.Len())
	}
	sl := l.ToSlice()
	for i := 0; i < len(s); i++ {
		if sl[i+3] != s[i] {
			t.Errorf("Expected item at index %d to be %v, but got %v", i+3, s[i], sl[i+3])
		}
	}
}

func TestLinkedListIndex(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 2, 2)
	i, err := l.Index(2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if i != 1 {
		t.Errorf("Expected index of 1, but got %d", i)
	}
}

func TestLinkedListInsert(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	l.Insert(1, 4)
	if l.Len() != 4 {
		t.Errorf("Expected length of 4, but got %d", l.Len())
	}
	sl := l.ToSlice()
	if sl[1] != 4 {
		t.Errorf("Expected item at index 1 to be 4, but got %v", sl[1])
	}
}

func TestLinkedListPop(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	x, err := l.Pop(1)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if x != 2 {
		t.Errorf("Expected popped item to be 2, but got %v", x)
	}
	if l.Len() != 2 {
		t.Errorf("Expected length of 2, but got %d", l.Len())
	}
}
func TestLinkedListRemove(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 2, 2)
	err := l.Remove(2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if l.Len() != 4 {
		t.Errorf("Expected length of 4, but got %d", l.Len())
	}
	sl := l.ToSlice()
	if sl[1] != 3 {
		t.Errorf("Expected item at index 1 to be 3, but got %v", sl[1])
	}
}
func TestLinkedListReverse(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	l.Reverse()
	if l.Len() != 3 {
		t.Errorf("Expected length of 3, but got %d", l.Len())
	}
	sl := l.ToSlice()
	if sl[0] != 3 || sl[1] != 2 || sl[2] != 1 {
		t.Errorf("Expected items to be [3, 2, 1], but got %v", sl)
	}
}
func TestLinkedListSortFunc(t *testing.T) {
	l := NewLinkedList(3, 1, 2, 0, 4, 8, 7, 9, 6, 5)
	l.SortFunc(func(a, b int) bool { return a < b })
	if l.Len() != 10 {
		t.Errorf("Expected length of 10, but got %d", l.Len())
	}
	sl := l.ToSlice()
	if !reflect.DeepEqual(sl, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("Expected items to be [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], but got %v", sl)
	}
}

func TestLinkedListSetItem(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.SetItem(1, 4)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if l.Len() != 3 {
		t.Errorf("Expected length of 3, but got %d", l.Len())
	}
	sl := l.ToSlice()
	if sl[1] != 4 {
		t.Errorf("Expected item at index 1 to be 4, but got %v", sl[1])
	}
	err = l.SetItem(-1, 20)
	if err != nil {
		t.Errorf("Unexpected error for Set(-1, 20): %v", err)
	}
	sl = l.ToSlice()
	if sl[2] != 20 {
		t.Errorf("Expected item at index 2 to be 20, but got %v", sl[2])
	}
}

func TestLinkedListGetItem(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	x, err := l.GetItem(1)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if x != 2 {
		t.Errorf("Expected item at index 1 to be 2, but got %v", x)
	}
	_, err = l.GetItem(6)
	if err == nil {
		t.Errorf("Expected error, but got %v", err)
	}
	x, err = l.GetItem(-1)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if x != 3 {
		t.Errorf("Expected item at index -1 to be 3, but got %v", x)
	}
}

func TestLinkedListLen(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	// Test Len
	if l.Len() != 3 {
		t.Errorf("Len failed. Expected 3, butgot %d", l.Len())
	}
}

// Test ToSlice and FromSlice
func TestLinkedListFromSliceToSlice(t *testing.T) {
	l := NewLinkedList[int]()
	s := []int{1, 2, 3, 4, 5}
	l.FromSlice(s)
	if !reflect.DeepEqual(l.ToSlice(), s) {
		t.Errorf("FromSlice and ToSlice failed. Expected %v, but got %v", s, l.ToSlice())
	}
}

func TestLinkedListSlice(t *testing.T) {
	l := NewLinkedList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Test start and default stop and step
	l1 := l.Slice(3)
	s := []int{3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}

	// Test start stop and default step
	l1 = l.Slice(3, 111)
	s = []int{3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}

	// Test nagtive start
	l1 = l.Slice(-2)
	s = []int{9, 10}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}

	// Test nagtive start less then -l.Len()
	l1 = l.Slice(-111)
	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}

	// Test nagtive start less then -l.Len()
	l1 = l.Slice(-111)
	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}

	// Test step * (stop - start) < 0
	l1 = l.Slice(1, 2, -1)
	s = []int{}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}

	l1 = l.Slice(2, 1, 2)
	s = []int{}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}

	// Test step < 0
	l1 = l.Slice(2, 0, -1)
	s = []int{2, 1}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}
	// Test stop < 0
	l1 = l.Slice(0, -1)
	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}

	// Test stop < 0 and step < 0
	l1 = l.Slice(-1, -6, -1)
	s = []int{10, 9, 8, 7, 6}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}

	// Test stop < -l.Len() and step < 0
	l1 = l.Slice(-1, -111, -1)
	s = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	if !reflect.DeepEqual(l1.ToSlice(), s) {
		t.Errorf("slice failed. Expected %v, but got %v", s, l1.ToSlice())
	}

	// Test step ==0
	assert.Panics(t, func() {
		l.Slice(1, 2, 0)
	}, "panic expected: slice step cannot be zero")
}

func TestLinkedListContains(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	// Test Contians
	if !l.Contains(1) || !l.Contains(2) || !l.Contains(3) {
		t.Errorf("Contains failed. Expected Contains 1, 2, 3, butgot not")
	}
}

func TestLinkedListEq(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	l1 := NewLinkedList(1, 2, 3)
	// Test Eq
	if !l.Eq(l1) {
		t.Errorf("Eq to LinkedList failed. ")
	}
	l2 := NewList(0, 1, 2, 3)
	if !l.Eq(l2) {
		t.Errorf("Eq to List failed. ")
	}
}
