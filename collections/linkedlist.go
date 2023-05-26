package collections

type Node[U comparable] struct {
	data U
	next *Node[U]
	pre  *Node[U]
}

func (n *Node[U]) Next() *Node[U] {
	return n.next
}

type LinkedList[U comparable] struct {
	head   *Node[U]
	tail   *Node[U]
	length int
}

func NewLinkedList[U comparable](items ...U) *LinkedList[U] {
	l := &LinkedList[U]{
		head: nil,
		tail: nil,
	}
	for _, v := range items {
		l.Append(v)
	}
	return l
}

func (l *LinkedList[U]) Head() *Node[U] {
	return l.head
}

func (l *LinkedList[U]) Append(item U) {
	n := &Node[U]{
		data: item,
	}
	if l.head != nil {
		n.pre = l.tail
		l.tail.next = n
		l.tail = n
	} else {
		l.head = n
		l.tail = n
	}
	l.length++
}

func (l *LinkedList[U]) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

func (l *LinkedList[U]) Copy() IList[U] {
	return NewLinkedList(l.ToSlice()...)
}

func (l *LinkedList[U]) Count(item U) int {
	cnt := 0
	for i := 0; i < l.Len(); i++ {
		v, _ := l.GetItem(i)
		if v == item {
			cnt += 1
		}
	}
	return cnt
}

func (l *LinkedList[U]) Extend(other any) error {
	switch val := other.(type) {
	case IList[U]:
		for i := 0; i < val.Len(); i++ {
			v, _ := val.GetItem(i)
			l.Append(v)
		}
	case []U:
		for _, v := range val {
			l.Append(v)
		}
	default:
		return ErrUnsupportType
	}
	return nil
}

// 超限返回nil
func (l *LinkedList[U]) getNodeFromIndex(index int) (n *Node[U]) {
	if l.head == nil {
		return nil
	}
	if index >= l.length || index < 0 {
		return nil
	}
	if index < l.length/2 {
		p := l.head
		for i := 0; p != nil && i <= index; i++ {
			if i == index {
				return p
			}
			p = p.next
		}
	} else {
		p := l.tail
		for i := l.length - 1; p != nil && i >= index; i-- {
			if i == index {
				return p
			}
			p = p.pre
		}
	}
	return
}

func (l *LinkedList[U]) Index(item U, args ...int) (int, error) {
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
	n := l.getNodeFromIndex(start)
	if n == nil {
		return -1, ErrElementNotFound
	}
	i := start
	for ; n != nil && i <= stop && n.data != item; i++ {
		n = n.next
	}
	if n != nil && n.data == item {
		return i, nil
	}
	return -1, ErrElementNotFound
}

func (l *LinkedList[U]) Insert(index int, item U) {
	if index+l.length < 0 {
		index = 0
	}
	if index >= l.Len() {
		l.Append(item)
		return
	}
	defer func() {
		l.length++
	}()
	n := l.getNodeFromIndex(index)
	nd := &Node[U]{
		data: item,
		pre:  n.pre,
	}
	n.pre = nd
	nd.next = n
	if nd.pre == nil {
		// n is head
		l.head = nd
	} else {
		nd.pre.next = nd
	}

}

func (l *LinkedList[U]) Pop(n ...int) (r U, err error) {
	index := l.Len() - 1
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
	defer func() {
		l.length--
	}()
	if index == 0 {
		r = l.head.data
		l.head = l.head.next
		if l.head != nil {
			l.head.pre = nil
		}
		return
	}
	if index == l.Len()-1 {
		r = l.tail.data
		l.tail = l.tail.pre
		if l.tail != nil {
			l.tail.next = nil
		}
		return
	}
	nd := l.getNodeFromIndex(index)
	r = nd.data
	nd.pre.next = nd.next
	nd.next.pre = nd.pre
	return
}

func (l *LinkedList[U]) getNodeFromData(item U) (n *Node[U]) {
	p := l.head
	for p != nil {
		if p.data == item {
			return p
		}
		p = p.next
	}
	return
}

func (l *LinkedList[U]) removeNode(n *Node[U]) {
	if n == l.head {
		l.head = l.head.next
		if l.head != nil {
			l.head.pre = nil
		}
	} else if n == l.tail {
		l.tail = l.tail.pre
		if l.tail != nil {
			l.tail.next = nil
		}
	} else {
		n.pre.next = n.next
		n.next.pre = n.pre
	}
}

func (l *LinkedList[U]) Remove(item U) error {
	n := l.getNodeFromData(item)
	if n == nil {
		return ErrElementNotFound
	}
	defer func() {
		l.length--
	}()
	l.removeNode(n)
	return nil
}

func (l *LinkedList[U]) Reverse() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		current.pre, current.next = current.next, current.pre
		current = current.pre
	}
	l.head, l.tail = l.tail, l.head
}

func (l *LinkedList[U]) SortFunc(less func(U, U) bool) {
	l.head = mergeSort(l.head, less)
}

func mergeSort[U comparable](head *Node[U], less func(U, U) bool) *Node[U] {
	if head == nil || head.next == nil {
		return head
	}

	slow, fast := head, head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	mid := slow.next
	slow.next = nil

	left := mergeSort(head, less)
	right := mergeSort(mid, less)

	return merge(left, right, less)
}

func merge[U comparable](left *Node[U], right *Node[U], less func(U, U) bool) *Node[U] {
	dummy := &Node[U]{}
	tail := dummy
	for left != nil && right != nil {
		if less(left.data, right.data) {
			tail.next = left
			left.pre = tail
			left = left.next
		} else {
			tail.next = right
			right.pre = tail
			right = right.next
		}
		tail = tail.next
	}

	if left != nil {
		tail.next = left
		left.pre = tail
	} else {
		tail.next = right
		right.pre = tail
	}
	return dummy.next
}

func (l *LinkedList[U]) SetItem(index int, item U) error {
	if index < 0 {
		index += l.length
	}
	n := l.getNodeFromIndex(index)
	if n == nil {
		return ErrElementNotFound
	}
	n.data = item
	return nil
}

func (l *LinkedList[U]) GetItem(index int) (r U, err error) {
	if index < 0 {
		index += l.length
	}
	n := l.getNodeFromIndex(index)
	if n == nil {
		return r, ErrElementNotFound
	}
	return n.data, nil
}

func (l *LinkedList[U]) Delete(index int) error {
	if index < 0 {
		index += l.length
	}
	if index < 0 || index > l.length {
		return ErrIndexOutofRange
	}
	_, err := l.Pop(index)
	return err
}

func (l *LinkedList[U]) Len() int {
	return l.length
}

func (l *LinkedList[U]) Slice(start int, args ...int) IList[U] {
	nl := NewLinkedList[U]()
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
		} else if step < 0 && stop < -1 {
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
	var p = l.getNodeFromIndex(start)
	if step > 0 {
		check = func(i, stop int) bool {
			return i < stop
		}
	} else {
		check = func(i, stop int) bool {
			return i > stop
		}
	}
	for i := start; check(i, stop) && p != nil; i += step {
		nl.Append(p.data)
		if step > 0 {
			for n := 0; n < step; n++ {
				p = p.next
			}
		} else {
			for n := 0; n < -step; n++ {
				p = p.pre
			}
		}
	}
	return nl
}

func (l *LinkedList[U]) ToSlice() []U {
	r := make([]U, 0, l.length)
	p := l.head
	for p != nil {
		r = append(r, p.data)
		p = p.next
	}
	return r
}

func (l *LinkedList[U]) FromSlice(items []U) {
	*l = LinkedList[U]{
		head: nil,
		tail: nil,
	}
	for _, v := range items {
		l.Append(v)
	}
}

func (l *LinkedList[U]) Contains(item U) bool {
	p := l.head
	for p != nil {
		if p.data == item {
			return true
		}
		p = p.next
	}
	return false
}

func (l *LinkedList[U]) Eq(other IList[U]) bool {
	if l.Len() != other.Len() {
		return false
	}
	p := l.head
	for i := 0; i < l.Len(); i++ {
		v, err := other.GetItem(i)
		if err != nil || p.data != v {
			return false
		}
		p = p.next
	}
	return true
}
func (l *LinkedList[U]) Swap(i, j int) {
	a := l.getNodeFromIndex(i)
	b := l.getNodeFromIndex(j)
	a.data, b.data = b.data, a.data
}
