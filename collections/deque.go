package collections

type Deque[T comparable] struct {
	left   IList[T]
	right  IList[T]
	MaxLen int
}

func NewDeque[T comparable](maxLen int, items ...T) *Deque[T] {
	if maxLen <= 0 {
		maxLen = -1
	}
	dq := &Deque[T]{
		left:   NewLinkedList[T](),
		right:  NewLinkedList[T](),
		MaxLen: maxLen,
	}
	dq.Extend(items)
	return dq
}

func (dq *Deque[T]) Len() int {
	return dq.left.Len() + dq.right.Len()
}

func (dq *Deque[T]) Append(t T) {
	if dq.Len() == dq.MaxLen {
		dq.PopLeft()
	}
	dq.right.Append(t)
}

func (dq *Deque[T]) AppendLeft(t T) {
	if dq.Len() == dq.MaxLen {
		dq.Pop()
	}
	dq.left.Append(t)
}

func (dq *Deque[T]) Clear() {
	dq.left.Clear()
	dq.right.Clear()
}

func (dq *Deque[T]) Copy() IList[T] {
	return &Deque[T]{
		left:   dq.left.Copy(),
		right:  dq.right.Copy(),
		MaxLen: dq.MaxLen,
	}
}

func (dq *Deque[T]) SortFunc(cmp func(T, T) bool) {
	panic("SortFunc is not implemented for deque")
}

func (dq *Deque[T]) Count(t T) int {
	return dq.left.Count(t) + dq.right.Count(t)
}

func (dq *Deque[T]) extendSlice(items []T) {
	for _, v := range items {
		dq.Append(v)
	}
}

func (dq *Deque[T]) extendLeftSlice(items []T) {
	for _, v := range items {
		dq.AppendLeft(v)
	}
}

func (dq *Deque[T]) extendDeque(other *Deque[T]) {
	for i := 0; i < other.Len(); i++ {
		v, _ := other.GetItem(i)
		dq.Append(v)
	}
}

func (dq *Deque[T]) extendLeftDeque(other *Deque[T]) {
	for i := 0; i < other.Len(); i++ {
		v, _ := other.GetItem(i)
		dq.AppendLeft(v)
	}
}

func (dq *Deque[T]) extendList(other IList[T]) {
	for i := 0; i < other.Len(); i++ {
		v, _ := other.GetItem(i)
		dq.Append(v)
	}
}

func (dq *Deque[T]) extendLeftList(other IList[T]) {
	for i := 0; i < other.Len(); i++ {
		v, _ := other.GetItem(i)
		dq.AppendLeft(v)
	}
}

func (dq *Deque[T]) Extend(other any) error {
	switch v := other.(type) {
	case []T:
		dq.extendSlice(v)
	case *Deque[T]:
		dq.extendDeque(v)
	case IList[T]:
		dq.extendList(v)
	default:
		return ErrUnsupportType
	}
	return nil
}

func (dq *Deque[T]) ExtendLeft(other any) error {
	switch v := other.(type) {
	case []T:
		dq.extendLeftSlice(v)
	case *Deque[T]:
		dq.extendLeftDeque(v)
	case IList[T]:
		dq.extendLeftList(v)
	default:
		return ErrUnsupportType
	}
	return nil
}

func (dq *Deque[T]) Index(t T, args ...int) (int, error) {
	var start = 0
	var stop = dq.Len()
	if len(args) > 0 {
		start = args[0]
	}
	if len(args) > 1 {
		stop = args[1]
	}
	for i := start; i < stop; i++ {
		v, _ := dq.GetItem(i)
		if v == t {
			return i, nil
		}
	}
	return -1, ErrElementNotFound
}

func (dq *Deque[T]) Insert(index int, t T) {
	if index < 0 {
		index += dq.Len()
	}
	if index < 0 {
		dq.AppendLeft(t)
		return
	}
	if index >= dq.Len() {
		dq.Append(t)
		return
	}
	if index < dq.left.Len() {
		dq.left.Insert(dq.left.Len()-index, t)
	} else {
		dq.right.Insert(index-dq.left.Len(), t)
	}
}

func (dq *Deque[T]) Pop(i ...int) (t T, err error) {
	if dq.right.Len() > 0 {
		return dq.right.Pop(-1)
	}
	return dq.left.Pop(0)
}

func (dq *Deque[T]) PopLeft() (t T, err error) {
	if dq.left.Len() > 0 {
		return dq.left.Pop(-1)
	}
	return dq.right.Pop(0)
}

func (dq *Deque[T]) Delete(index int) error {
	if index < 0 {
		index += dq.Len()
	}
	if index < dq.left.Len() {
		return dq.left.Delete(dq.left.Len() - index - 1)
	} else {
		return dq.right.Delete(index - dq.left.Len())
	}
}

func (dq *Deque[T]) Remove(t T) (err error) {
	for i := 0; i < dq.Len(); i++ {
		v, _ := dq.GetItem(i)
		if v == t {
			return dq.Delete(i)
		}
	}
	return ErrElementNotFound
}

func (dq *Deque[T]) Reverse() {
	dq.left, dq.right = dq.right, dq.left
}

func (dq *Deque[T]) Rotate(n ...int) {
	if dq.Len() == 0 {
		return
	}
	num := 1
	if len(n) > 0 {
		num = n[0]
	}
	step := 1
	if num < 0 {
		step = -1
	}
	for ; num != 0; num -= step {
		if num > 0 {
			v, _ := dq.Pop()
			dq.AppendLeft(v)
		} else {
			v, _ := dq.PopLeft()
			dq.Append(v)
		}
	}
}

func (dq *Deque[T]) GetItem(index int) (r T, err error) {
	if index < 0 {
		index += dq.Len()
	}
	if index < 0 {
		err = ErrIndexOutofRange
		return
	}
	if index >= dq.Len() {
		err = ErrIndexOutofRange
		return
	}
	if index < dq.left.Len() {
		return dq.left.GetItem(index)
	}
	return dq.right.GetItem(index - dq.left.Len())
}

func (dq *Deque[T]) SetItem(index int, value T) (err error) {
	if index < 0 {
		index += dq.Len()
	}
	if index < 0 {
		err = ErrIndexOutofRange
		return
	}
	if index >= dq.Len() {
		err = ErrIndexOutofRange
		return
	}
	if index < dq.left.Len() {
		return dq.left.SetItem(index, value)
	}
	return dq.right.SetItem(index-dq.left.Len(), value)
}

func (dq *Deque[T]) Contains(item T) bool {
	return dq.left.Contains(item) || dq.right.Contains(item)
}

func (dq *Deque[T]) Slice(start int, args ...int) IList[T] {
	panic("Slice is not implemented for deque")
}

func (dq *Deque[T]) ToSlice() []T {
	l := dq.left.Copy()
	l.Reverse()
	return append(l.ToSlice(), dq.right.ToSlice()...)
}

func (dq *Deque[T]) FromSlice(slice []T) {
	*dq = *NewDeque(dq.MaxLen, slice...)
}

func (dq *Deque[T]) Eq(other IList[T]) bool {
	if dq.Len() != other.Len() {
		return false
	}
	for i := 0; i < dq.Len(); i++ {
		v1, _ := dq.GetItem(i)
		v2, _ := other.GetItem(i)
		if v1 != v2 {
			return false
		}
	}
	return true
}
