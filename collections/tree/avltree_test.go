package tree

import (
	"reflect"
	"testing"
)

func TestAVLTree_NewAVLTreeByKey(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}
	tree := NewAVLTreeByKey(func(u *user) int { return u.ID })
	tree.Insert(&user{5, "alice"})
	tree.Insert(&user{2, "bob"})
	tree.Insert(&user{7, "tom"})
	tree.Insert(&user{1, "sam"})
	tree.Insert(&user{8, "grace"})
	tree.Insert(&user{3, "lily"})
	tree.Insert(&user{6, "jim"})
	tree.Insert(&user{4, "candy"})
	tree.Insert(&user{5, "bob"})
	if !tree.IsBalanced() {
		t.Errorf("tree insert Failed, not balanced")
	}
	users := make([]*user, 0)
	tree.Traverse(func(u *user) { users = append(users, u) })
	if len(users) != 8 {
		t.Errorf("tree insert Failed, length error")
	}
	u, err := tree.Search(&user{ID: 4})
	if err != nil {
		t.Errorf("Unexpected error when searching for an existing element: %v", err)
	}
	if u.Name != "candy" {
		t.Errorf("Expected to find element candy, but found %v", u.Name)
	}
	_, err = tree.Search(&user{ID: 10})
	if err == nil {
		t.Errorf("Expected an error when searching for a non-existent element")
	}
}

func TestAVLTree_Insert(t *testing.T) {
	tree := NewAVLTree(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	})
	if !tree.IsBalanced() {
		t.Errorf("tree init Failed, new tree is not balanced")
	}
	tree.Insert(5)
	if !tree.IsBalanced() {
		t.Errorf("tree insert Failed, not balanced")
	}
	tree.Insert(2)
	if !tree.IsBalanced() {
		t.Errorf("tree insert Failed, not balanced")
	}
	tree.Insert(7)
	if !tree.IsBalanced() {
		t.Errorf("tree insert Failed, not balanced")
	}
	tree.Insert(1)
	if !tree.IsBalanced() {
		t.Errorf("tree insert Failed, not balanced")
	}
	tree.Insert(8)
	if !tree.IsBalanced() {
		t.Errorf("tree insert Failed, not balanced")
	}
	tree.Insert(3)
	if !tree.IsBalanced() {
		t.Errorf("tree insert Failed, not balanced")
	}
	tree.Insert(6)
	if !tree.IsBalanced() {
		t.Errorf("tree insert Failed, not balanced")
	}
	tree.Insert(4)
	if !tree.IsBalanced() {
		t.Errorf("tree insert Failed, not balanced")
	}
	tree.Insert(5) // 重复元素不应该被插入
	if !tree.IsBalanced() {
		t.Errorf("tree insert Failed, not balanced")
	}
	// 检查树中每个节点高度是否设置正确
	nodes := make([]*AVLNode[int], 0)
	tree.TraverseNode(func(v *AVLNode[int]) {
		if v.height != tree.height(v) {
			nodes = append(nodes, v)
		}
	})
	if len(nodes) != 0 {
		t.Errorf("tree height error")
	}
	// 检查树是否正确构建
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	actual := make([]int, 0, len(expected))
	tree.Traverse(func(v int) { actual = append(actual, v) })
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected tree structure. Expected %v, but got %v", expected, actual)
	}
}

func TestAVLTree_Delete(t *testing.T) {
	tree := NewAVLTree(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	})

	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(4)

	// 删除不存在的元素，应返回错误
	err := tree.Delete(10)
	if err == nil {
		t.Errorf("Expected an error when deleting a non-existent element")
	}

	if tree.root.value != 5 {
		t.Errorf("root is not 5")
	}
	// 删除现有元素
	err = tree.Delete(5)
	if err != nil {
		t.Errorf("Unexpected error when deleting an existing element: %v", err)
	}

	if tree.root.value != 6 {
		t.Errorf("root is not 6, is %v", tree.root.value)
	}
	// 检查树是否正确重构
	expected := []int{1, 2, 3, 4, 6, 7, 8}
	actual := make([]int, 0, len(expected))
	tree.Traverse(func(v int) { actual = append(actual, v) })
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected tree structure. Expected %v, but got %v", expected, actual)
	}
}

func TestAVLTree_Search(t *testing.T) {
	tree := NewAVLTree(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	})

	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(4)

	// 搜索存在的元素
	val, err := tree.Search(6)
	if err != nil {
		t.Errorf("Unexpected error when searching for an existing element: %v", err)
	}
	if val != 6 {
		t.Errorf("Expected to find element 6, but found %v", val)
	}

	// 搜索不存在的元素，应返回错误
	_, err = tree.Search(10)
	if err == nil {
		t.Errorf("Expected an error when searching for a non-existent element")
	}
}
