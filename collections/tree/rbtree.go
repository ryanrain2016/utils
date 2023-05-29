package tree

import (
	"math"

	"github.com/ryanrain2016/utils/collections"
	"github.com/ryanrain2016/utils/types"
)

type RBNode[T any] struct {
	value T
	left  *RBNode[T]
	right *RBNode[T]
	red   bool
}

type RBTree[T any] struct {
	root *RBNode[T]
	size int
	cmp  func(T, T) int
}

func NewRBTree[T any](cmp func(T, T) int) *RBTree[T] {
	return &RBTree[T]{nil, 0, cmp}
}

func NewRBTreeByKey[T any, U types.Ordered](key func(T) U) *RBTree[T] {
	return &RBTree[T]{nil, 0, func(t1, t2 T) int {
		k1, k2 := key(t1), key(t2)
		if k1 < k2 {
			return -1
		} else if k1 == k2 {
			return 0
		}
		return 1
	}}
}

func (t *RBTree[T]) insert(node *RBNode[T], value T) (n *RBNode[T], ok bool) {
	if node == nil {
		return &RBNode[T]{value, nil, nil, true}, true
	}
	r := t.cmp(value, node.value)
	if r < 0 {
		node.left, ok = t.insert(node.left, value)
	} else if r > 0 {
		node.right, ok = t.insert(node.right, value)
	} else {
		node.value = value
		return node, false
	}

	if isRed(node.right) && !isRed(node.left) {
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		flipColors(node)
	}

	return node, ok
}

func (t *RBTree[T]) Insert(value T) {
	root, ok := t.insert(t.root, value)
	t.root = root
	t.root.red = false
	if ok {
		t.size++
	}
}

func (t *RBTree[T]) remove(node *RBNode[T], value T) (*RBNode[T], error) {
	if node == nil {
		return nil, collections.ErrElementNotFound
	}
	r := t.cmp(value, node.value)
	if r < 0 {
		left, err := t.remove(node.left, value)
		if err != nil {
			return nil, err
		}
		node.left = left
	} else if r > 0 {
		right, err := t.remove(node.right, value)
		if err != nil {
			return nil, err
		}
		node.right = right
	} else {
		if node.left == nil {
			return node.right, nil
		} else if node.right == nil {
			return node.left, nil
		} else {
			smallest := findSmallest(node.right)
			node.value = smallest.value
			node.right = deleteMin(node.right)
		}
	}

	return fixUp(node), nil
}

func (t *RBTree[T]) Contains(value T) bool {
	if node := t.search(t.root, value); node != nil {
		return true
	}
	return false
}

func (t *RBTree[T]) Remove(value T) error {
	root, err := t.remove(t.root, value)
	if err != nil {
		return err
	}
	t.root = root
	t.size--
	if t.root != nil {
		t.root.red = false
	}
	return nil
}

func (t *RBTree[T]) search(node *RBNode[T], value T) *RBNode[T] {
	if node == nil {
		return nil
	}
	r := t.cmp(value, node.value)
	if r < 0 {
		return t.search(node.left, value)
	} else if r > 0 {
		return t.search(node.right, value)
	} else {
		return node
	}
}

func (t *RBTree[T]) Search(value T) (r T, err error) {
	node := t.search(t.root, value)
	if node == nil {
		return r, collections.ErrElementNotFound
	}
	return node.value, nil
}

func (t *RBTree[T]) traverse(node *RBNode[T], f func(*RBNode[T])) {
	if node == nil {
		f(node)
		return
	}

	t.traverse(node.left, f)
	f(node)
	t.traverse(node.right, f)
}

func (t *RBTree[T]) Traverse(f func(T)) {
	t.traverse(t.root, func(r *RBNode[T]) {
		if r == nil {
			return
		}
		f(r.value)
	})
}

func isRed[T any](node *RBNode[T]) bool {
	if node == nil {
		return false
	}
	return node.red
}

func isBlack[T any](n *RBNode[T]) bool {
	return n == nil || !n.red
}

func rotateLeft[T any](node *RBNode[T]) *RBNode[T] {
	x := node.right
	node.right = x.left
	x.left = node
	x.red = node.red
	node.red = true
	return x
}

func rotateRight[T any](node *RBNode[T]) *RBNode[T] {
	x := node.left
	node.left = x.right
	x.right = node
	x.red = node.red
	node.red = true
	return x
}

func flipColors[T any](node *RBNode[T]) {
	node.red = !node.red
	node.left.red = !node.left.red
	node.right.red = !node.right.red
}

func moveRedLeft[T any](node *RBNode[T]) *RBNode[T] {
	flipColors(node)
	if isRed(node.right.left) {
		node.right = rotateRight(node.right)
		node = rotateLeft(node)
		flipColors(node)
	}
	return node
}

func moveRedRight[T any](node *RBNode[T]) *RBNode[T] {
	flipColors(node)
	if isRed(node.left.left) {
		node = rotateRight(node)
		flipColors(node)
	}
	return node
}

func findSmallest[T any](node *RBNode[T]) *RBNode[T] {
	if node.left == nil {
		return node
	}
	return findSmallest(node.left)
}

func deleteMin[T any](node *RBNode[T]) *RBNode[T] {
	if node.left == nil {
		return nil
	}
	if !isRed(node.left) && !isRed(node.left.left) {
		node = moveRedLeft(node)
	}
	node.left = deleteMin(node.left)
	return fixUp(node)
}

func fixUp[T any](node *RBNode[T]) *RBNode[T] {
	if isRed(node.right) {
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		flipColors(node)
	}
	return node
}

func (t *RBTree[T]) Minimum() *RBNode[T] {
	node := t.root
	for node != nil && node.left != nil {
		node = node.left
	}
	return node
}

func (t *RBTree[T]) Maximum() *RBNode[T] {
	node := t.root
	for node != nil && node.right != nil {
		node = node.right
	}
	return node
}

func (t *RBTree[T]) Size() int {
	return t.size
}

// Height returns the height of the tree.
func (t *RBTree[T]) Height() int {
	return t.height(t.root)
}

func (t *RBTree[T]) height(node *RBNode[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := t.height(node.left)
	rightHeight := t.height(node.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func (t *RBTree[T]) inOrder(node *RBNode[T], f func(*RBNode[T])) {
	f(node)
	if node == nil {
		return
	}
	t.inOrder(node.left, f)
	t.inOrder(node.right, f)
}

func (t *RBTree[T]) InOrder(f func(T)) {
	t.inOrder(t.root, func(r *RBNode[T]) {
		f(r.value)
	})
}

func (t *RBTree[T]) toMap(node *RBNode[T]) map[string]any {
	if node == nil {
		return nil
	}
	r := make(map[string]any)
	r["value"] = node.value
	r["left"] = t.toMap(node.left)
	r["right"] = t.toMap(node.right)
	if node.red {
		r["color"] = "red"
	} else {
		r["color"] = "black"
	}

	return r
}

// for debug
func (t *RBTree[T]) ToMap() map[string]any {
	return t.toMap(t.root)
}

// for debug and test
func (t *RBTree[T]) checkBlackNodeForEveryPath(n *RBNode[T]) bool {
	if n == nil {
		return true
	}
	_, err := t.blackNodesCount(t.root)
	return err == nil
}

func (t *RBTree[T]) blackNodesCount(n *RBNode[T]) (int, error) {
	if n == nil {
		return 0, nil
	}
	left, err := t.blackNodesCount(n.left)
	if err != nil {
		return 0, err
	}
	right, err := t.blackNodesCount(n.right)
	if err != nil {
		return 0, err
	}
	if left != right {
		return 0, err
	}
	if isBlack(n) {
		return left + 1, nil
	} else {
		return left, nil
	}
}

// for debug and test
func (t *RBTree[T]) isValid() bool {
	// 根节点是黑色的
	if t.root != nil && t.root.red {
		return false
	}
	cnt := 0
	// 如果一个节点是红色的，则它的两个子节点都是黑色的
	t.traverse(t.root, func(r *RBNode[T]) {
		if cnt != 0 {
			return
		}
		if !(r == nil || (!r.red) || (r.red && (r.left == nil || !r.left.red) && (r.right == nil || !r.right.red))) {
			cnt += 1
		}
	})
	if cnt != 0 {
		return false
	}
	cnt = 0
	// 平衡
	t.traverse(t.root, func(r *RBNode[T]) {
		if cnt != 0 {
			return
		}
		if r == nil {
			return
		}
		sub := t.height(r.left) - t.height(r.right)
		if sub < -1 || sub > 1 {
			cnt++
		}
	})
	cnt = 0
	// 对于每个节点，从该节点到其所有后代叶子节点的简单路径上，均包含相同数目的黑色节点
	// 高度
	return t.checkBlackNodeForEveryPath(t.root) && t.Height() <= int(2*math.Log2(float64(t.Size()+1)))
}
