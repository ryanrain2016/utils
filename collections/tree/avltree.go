package tree

import (
	"github.com/ryanrain2016/utils/collections"
	"github.com/ryanrain2016/utils/types"
)

type AVLTree[T any] struct {
	root *AVLNode[T]
	cmp  func(T, T) int
}

type AVLNode[T any] struct {
	value  T
	left   *AVLNode[T]
	right  *AVLNode[T]
	height int
}

func NewAVLTree[T any](cmp func(T, T) int) *AVLTree[T] {
	return &AVLTree[T]{cmp: cmp}
}

func NewAVLTreeByKey[T any, U types.Ordered](key func(T) U) *AVLTree[T] {
	return &AVLTree[T]{cmp: func(t1, t2 T) int {
		k1, k2 := key(t1), key(t2)
		if k1 < k2 {
			return -1
		} else if k1 > k2 {
			return 1
		} else {
			return 0
		}
	}}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func (t *AVLTree[T]) insert(root *AVLNode[T], value T) *AVLNode[T] {
	if root == nil {
		return &AVLNode[T]{value: value, height: 1}
	}

	if t.cmp(value, root.value) < 0 {
		root.left = t.insert(root.left, value)
	} else if t.cmp(value, root.value) > 0 {
		root.right = t.insert(root.right, value)
	} else {
		root.value = value
		return root
	}

	root.height = 1 + max(t.height(root.left), t.height(root.right))

	balance := t.getBalance(root)

	if balance > 1 && t.cmp(value, root.left.value) < 0 {
		return t.rightRotate(root)
	}

	if balance < -1 && t.cmp(value, root.right.value) > 0 {
		return t.leftRotate(root)
	}

	if balance > 1 && t.cmp(value, root.left.value) > 0 {
		root.left = t.leftRotate(root.left)
		return t.rightRotate(root)
	}

	if balance < -1 && t.cmp(value, root.right.value) < 0 {
		root.right = t.rightRotate(root.right)
		return t.leftRotate(root)
	}

	return root
}

func (t *AVLTree[T]) Insert(value T) {
	t.root = t.insert(t.root, value)
}

func (t *AVLTree[T]) deleteAVLNode(root *AVLNode[T], value T) (n *AVLNode[T], err error) {
	if root == nil {
		return nil, collections.ErrElementNotFound
	}
	switch t.cmp(value, root.value) {
	case -1:
		left, err := t.deleteAVLNode(root.left, value)
		if err != nil {
			return nil, err
		}
		root.left = left
	case 1:
		right, err := t.deleteAVLNode(root.right, value)
		if err != nil {
			return nil, err
		}
		root.right = right
	default:
		if root.left == nil || root.right == nil {
			var temp *AVLNode[T]
			if root.left != nil {
				temp = root.left
			} else {
				temp = root.right
			}

			if temp == nil {
				root = nil
			} else {
				*root = *temp
			}
		} else {
			temp := t.minValueAVLNode(root.right)
			root.value = temp.value
			root.right, _ = t.deleteAVLNode(root.right, temp.value)
		}
	}

	if root == nil {
		return nil, nil
	}

	root.height = 1 + max(t.height(root.left), t.height(root.right))

	balance := t.getBalance(root)

	if balance > 1 && t.getBalance(root.left) >= 0 {
		return t.rightRotate(root), nil
	}

	if balance < -1 && t.getBalance(root.right) <= 0 {
		return t.leftRotate(root), nil
	}

	if balance > 1 && t.getBalance(root.left) < 0 {
		root.left = t.leftRotate(root.left)
		return t.rightRotate(root), nil
	}

	if balance < -1 && t.getBalance(root.right) > 0 {
		root.right = t.rightRotate(root.right)
		return t.leftRotate(root), nil
	}

	return root, nil
}

func (t *AVLTree[T]) Delete(value T) error {
	var err error
	root, err := t.deleteAVLNode(t.root, value)
	if err == nil {
		t.root = root
	}
	return err
}

func (t *AVLTree[T]) height(root *AVLNode[T]) int {
	if root == nil {
		return 0
	}
	return root.height
}

func (t *AVLTree[T]) getBalance(root *AVLNode[T]) int {
	if root == nil {
		return 0
	}
	return t.height(root.left) - t.height(root.right)
}

func (t *AVLTree[T]) minValueAVLNode(node *AVLNode[T]) *AVLNode[T] {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (t *AVLTree[T]) leftRotate(x *AVLNode[T]) *AVLNode[T] {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = 1 + max(t.height(x.left), t.height(x.right))
	y.height = 1 + max(t.height(y.left), t.height(y.right))

	return y
}

func (t *AVLTree[T]) rightRotate(y *AVLNode[T]) *AVLNode[T] {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = 1 + max(t.height(y.left), t.height(y.right))
	x.height = 1 + max(t.height(x.left), t.height(x.right))

	return x
}

func (t *AVLTree[T]) Traverse(f func(T)) {
	t.traverse(t.root, f)
}

func (t *AVLTree[T]) traverse(node *AVLNode[T], f func(T)) {
	if node != nil {
		t.traverse(node.left, f)
		f(node.value)
		t.traverse(node.right, f)
	}
}

func (t *AVLTree[T]) TraverseNode(f func(*AVLNode[T])) {
	t.traverseNode(t.root, f)
}

func (t *AVLTree[T]) traverseNode(node *AVLNode[T], f func(*AVLNode[T])) {
	if node != nil {
		t.traverseNode(node.left, f)
		f(node)
		t.traverseNode(node.right, f)
	}
}

func (t *AVLTree[T]) IsBalanced() bool {
	return t.isBalanced(t.root)
}

func (t *AVLTree[T]) isBalanced(node *AVLNode[T]) bool {
	if node == nil {
		return true
	}
	leftHeight := t.height(node.left)
	rightHeight := t.height(node.right)

	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	return t.isBalanced(node.left) && t.isBalanced(node.right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (t *AVLTree[T]) Search(target T) (r T, err error) {
	node := t.search(t.root, target)
	if node == nil {
		err = collections.ErrElementNotFound
		return
	}
	return node.value, nil
}

func (t *AVLTree[T]) search(node *AVLNode[T], target T) *AVLNode[T] {
	if node == nil {
		return nil
	}
	if t.cmp(target, node.value) == 0 {
		return node
	} else if t.cmp(target, node.value) < 0 {
		return t.search(node.left, target)
	} else {
		return t.search(node.right, target)
	}
}
