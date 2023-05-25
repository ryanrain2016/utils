package collections

import (
	"fmt"
	"sort"

	"github.com/ryanrain2016/utils/itertools"
)

type IList[U comparable] interface {
	Append(item U)
	Clear()
	Copy() IList[U]
	Count(item U) int
	Extend(any) error
	Index(U, ...int) (int, error)
	Insert(index int, item U)
	Pop(index ...int) (U, error)
	Remove(item U) error
	Reverse()
	SortFunc(less func(U, U) bool)
	SetItem(index int, item U) error
	GetItem(index int) (U, error)
	Delete(index int) error
	Len() int
	Slice(int, ...int) IList[U]
	ToSlice() []U
	FromSlice([]U)
	Contains(U) bool
	Eq(other IList[U]) bool
}

type List[U comparable] struct {
	items []U
}

func NewList[U comparable](size int, items ...U) *List[U] {
	l := len(items)
	if size < l {
		size = l
	}
	items_ := make([]U, l, size)
	copy(items_, items)
	return &List[U]{
		items: items_,
	}
}

func (l *List[U]) Append(item U) {
	l.items = append(l.items, item)
}

func (l *List[U]) Clear() {
	l.items = make([]U, 0, cap(l.items))
}

func (l *List[U]) Copy() IList[U] {
	items := make([]U, len(l.items), cap(l.items))
	copy(items, l.items)
	return &List[U]{
		items: items,
	}
}

func (l *List[U]) Count(item U) int {
	cnt := 0
	for i := 0; i < l.Len(); i++ {
		if l.items[i] == item {
			cnt++
		}
	}
	return cnt
}

func (l *List[U]) Extend(other any) error {
	switch val := other.(type) {
	case IList[U]:
		for i := 0; i < val.Len(); i++ {
			v, _ := val.GetItem(i)
			l.Append(v)
		}
	case []U:
		l.items = append(l.items, val...)
	default:
		return ErrUnsupportType
	}
	return nil
}

func (l *List[U]) Index(item U, args ...int) (int, error) {
	var start = 0
	var stop = l.Len()
	if len(args) > 0 {
		start = args[0]
	}
	if start < 0 {
		start += l.Len()
	}
	if len(args) > 1 {
		stop = args[1]
	}
	if stop < 0 {
		stop += l.Len()
	}

	for i := start; i < stop; i++ {
		if l.items[i] == item {
			return i, nil
		}
	}
	return -1, fmt.Errorf("%v is not in list", item)
}

func (l *List[U]) Insert(index int, item U) {
	if index < 0 {
		index += len(l.items)
	}
	if index < 0 {
		index = 0
	}
	l.items = append(l.items, item)
	for i := len(l.items) - 1; i > index; i-- {
		l.items[i] = l.items[i-1]
	}
	if index < len(l.items) {
		l.items[index] = item
	}
}

func (l *List[U]) Pop(n ...int) (item U, err error) {
	index := -1
	if len(n) > 0 {
		index = n[0]
	}
	if index < 0 {
		index += l.Len()
	}
	if index < 0 || index >= l.Len() {
		err = ErrIndexOutofRange
		return
	}
	item = l.items[index]
	copy(l.items[index:], l.items[index+1:])
	l.items = l.items[:len(l.items)-1]
	return
}

func (l *List[U]) Remove(item U) error {
	for i := 0; i < l.Len(); i++ {
		if l.items[i] == item {
			l.Pop(i)
			return nil
		}
	}
	return ErrElementNotFound
}

func (l *List[U]) Reverse() {
	itertools.ReverseSlice(l.items)
}

func (l *List[U]) SortFunc(less func(U, U) bool) {
	sort.Slice(l.items, func(i, j int) bool {
		return less(l.items[i], l.items[j])
	})
}

func (l *List[U]) SetItem(index int, item U) (err error) {
	if index < 0 {
		index += l.Len()
	}
	if index < 0 {
		err = ErrIndexOutofRange
		return
	}
	if index >= l.Len() {
		err = ErrIndexOutofRange
		return
	}
	l.items[index] = item
	return nil
}

func (l *List[U]) GetItem(index int) (r U, err error) {
	if index < 0 {
		index += l.Len()
	}
	if index < 0 {
		err = ErrIndexOutofRange
		return
	}
	if index >= l.Len() {
		err = ErrIndexOutofRange
		return
	}
	return l.items[index], nil
}

func (l *List[U]) Delete(index int) error {
	if _, err := l.Pop(index); err != nil {
		return ErrIndexOutofRange
	}
	return nil
}

func (l *List[U]) Len() int {
	return len(l.items)
}

func (l *List[U]) Slice(start int, args ...int) IList[U] {
	nl := NewList[U](0)
	step := 1
	if len(args) > 1 {
		step = args[1]
	}
	if step == 0 {
		panic("slice step cannot be zero")
	}

	if start < 0 {
		start += l.Len()
	}
	if start > l.Len()-1 {
		if step < 0 {
			start = l.Len() - 1
		} else {
			return nl
		}
	}
	if start < 0 {
		if step > 0 {
			start = 0
		} else {
			return nl
		}
	}
	var stop int
	if len(args) > 0 {
		stop = args[0]
		if stop < 0 {
			stop += l.Len()
		}
		if step > 0 && stop > l.Len() {
			stop = l.Len()
		}
		if step < 0 && stop < -1 {
			stop = -1
		}
	} else {
		if step > 0 {
			stop = l.Len()
		} else {
			stop = -1
		}
	}

	var check func(i, stop int) bool
	if step > 0 {
		check = func(i, stop int) bool {
			return i < stop
		}
	} else {
		check = func(i, stop int) bool {
			return i > stop
		}
	}
	for i := start; check(i, stop); i += step {
		item, _ := l.GetItem(i)
		nl.Append(item)
	}
	return nl
}

func (l *List[U]) ToSlice() []U {
	items := make([]U, l.Len())
	copy(items, l.items)
	return items
}

func (l *List[U]) FromSlice(slice []U) {
	*l = *NewList(0, slice...)
}

func (l *List[U]) Contains(item U) bool {
	for _, v := range l.items {
		if v == item {
			return true
		}
	}
	return false
}

func (l *List[U]) Eq(other IList[U]) bool {
	if v, ok := other.(*LinkedList[U]); ok {
		// LinkedList.GetItem is not efficient
		return v.Eq(l)
	}
	if l.Len() != other.Len() {
		return false
	}
	for i := 0; i < l.Len(); i++ {
		v1, _ := l.GetItem(i)
		v2, _ := other.GetItem(i)
		if v1 != v2 {
			return false
		}
	}
	return true
}

func (l *List[U]) Swap(i, j int) {
	l.items[i], l.items[j] = l.items[j], l.items[i]
}
