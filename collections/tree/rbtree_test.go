package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func cmp(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}

func TestInsert(t *testing.T) {
	tree := NewRBTree(cmp)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(23)
	tree.Insert(28)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	if !tree.isValid() {
		t.Errorf("tree is not valid after insert")
	}
}

func TestInsertDuplicate(t *testing.T) {
	tree := NewRBTree(cmp)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(4)
	tree.Insert(5)

	if !tree.isValid() {
		t.Errorf("tree is not valid after insert")
	}
	if tree.Size() != 8 {
		t.Errorf("tree insert duplicate element should not add")
	}
}

func TestDelete(t *testing.T) {
	tree := NewRBTree(cmp)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(23)
	tree.Insert(28)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	// 删除元素 15、25、10
	tree.Remove(15)
	tree.Remove(25)
	tree.Remove(10)

	if !tree.isValid() {
		t.Errorf("tree is not valid after remove")
	}

	var result []int
	tree.Traverse(func(i int) {
		result = append(result, i)
	})

	if !reflect.DeepEqual(result, []int{3, 5, 7, 20, 23, 28, 30}) {
		t.Errorf("Traverse failed")
	}
}

func TestSearchExists(t *testing.T) {
	tree := NewRBTree(cmp)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(23)
	tree.Insert(28)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	if !tree.isValid() {
		t.Errorf("tree is not valid after insert")
	}
	// 检查元素 20 是否存在
	if r, err := tree.Search(20); err != nil {
		t.Errorf("Element 20 should exist in the tree")
	} else if r != 20 {
		t.Errorf("Search Failed expect 20, but got %v", r)
	}

	if !tree.isValid() {
		t.Errorf("tree is not valid after search")
	}
}

func TestSearchNotExists(t *testing.T) {
	tree := NewRBTree(cmp)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(23)
	tree.Insert(28)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	// 检查元素 100 是否存在
	if _, err := tree.Search(100); err == nil {
		t.Errorf("Element 100 should not exist in the tree")
	}

	if !tree.isValid() {
		t.Errorf("tree is not valid after search")
	}
}

func TestTraversal(t *testing.T) {
	tree := NewRBTree(cmp)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(23)
	tree.Insert(28)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	// 中序遍历
	var result []int
	tree.Traverse(func(i int) {
		result = append(result, i)
	})

	if !reflect.DeepEqual(result, []int{3, 5, 7, 10, 15, 20, 23, 25, 28, 30}) {
		t.Errorf("Traverse failed")
	}
	if !tree.isValid() {
		t.Errorf("tree is not valid after traverse")
	}
}

func TestDeleteNotExists(t *testing.T) {
	tree := NewRBTree(cmp)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(23)
	tree.Insert(28)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	if !tree.isValid() {
		t.Errorf("tree is not valid after insert")
	}
	// 删除元素 100
	err := tree.Remove(100)
	if err == nil {
		t.Errorf("Deleting non-existent element should return an error")
	}

	// 删除元素 21
	err = tree.Remove(21)
	if err == nil {
		t.Errorf("Deleting non-existent element should return an error")
	}

	// 删除元素 13
	err = tree.Remove(13)
	if err == nil {
		t.Errorf("Deleting non-existent element should return an error")
	}

	if !tree.isValid() {
		t.Errorf("tree is not valid after remove")
		fmt.Printf("tree: %v\n", tree.Size())
	}
}

func TestDeleteExists(t *testing.T) {
	tree := NewRBTree(cmp)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(23)
	tree.Insert(28)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	// 删除元素 20
	err := tree.Remove(20)
	if err != nil {
		t.Errorf("Deleting existing element should not return an error")
	}

	if !tree.isValid() {
		t.Errorf("tree is not valid after remove")
	}

	// 删除元素 3
	err = tree.Remove(3)
	if err != nil {
		t.Errorf("Deleting existing element should not return an error")
	}

	if !tree.isValid() {
		t.Errorf("tree is not valid after remove")
	}
}

func TestNewRBTreeByKey(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}
	tree := NewRBTreeByKey(func(u *user) int { return u.ID })
	tree.Insert(&user{5, "alice"})
	tree.Insert(&user{2, "bob"})
	tree.Insert(&user{7, "tom"})
	tree.Insert(&user{1, "sam"})
	tree.Insert(&user{8, "grace"})
	tree.Insert(&user{3, "lily"})
	tree.Insert(&user{6, "jim"})
	tree.Insert(&user{4, "candy"})
	tree.Insert(&user{5, "bob"})

	if !tree.isValid() {
		t.Errorf("tree is not valid after insert")
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
