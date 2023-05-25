package collections

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderedDict(t *testing.T) {
	dict := NewOrderedDict[int, string]([]*MapItem[int, string]{{1, "one"}, {2, "two"}})
	if len(dict.Keys()) != 2 {
		t.Errorf("Expected length of keys to be 2, got %d", len(dict.Keys()))
	}
}

func TestOrderedDict_Update(t *testing.T) {
	dict := NewOrderedDict[int, string]()
	dict.Update(map[int]string{1: "one", 2: "two"})
	if len(dict.Keys()) != 2 {
		t.Errorf("Expected length of keys to be 2, got %d", len(dict.Keys()))
	}
	v, err := dict.GetItem(1)
	assert.Nil(t, err)
	if v != "one" {
		t.Errorf("Expected value for key 1 to be 'one', got '%s'", v)
	}
}

func TestOrderedDict_UpdatePairSlice(t *testing.T) {
	dict := NewOrderedDict[int, string]()
	dict.UpdatePairSlice([]int{1, 2}, []string{"one", "two"})
	if len(dict.Keys()) != 2 {
		t.Errorf("Expected length of keys to be 2, got %d", len(dict.Keys()))
	}
	v, err := dict.GetItem(1)
	assert.Nil(t, err)
	if v != "one" {
		t.Errorf("Expected value for key 1 to be 'one', got '%s'", v)
	}
}

func TestOrderedDict_Items(t *testing.T) {
	dict := NewOrderedDict[int, string]([]*MapItem[int, string]{{1, "one"}, {2, "two"}})
	items := dict.Items()
	if len(items) != 2 {
		t.Errorf("Expected length of items to be 2, got %d", len(items))
	}
	if items[0].Key != 1 || items[0].Value != "one" {
		t.Errorf("Expected item 0 to be {1, 'one'}, got %v", items[0])
	}
}

func TestOrderedDict_PopItem(t *testing.T) {
	dict := NewOrderedDict[int, string]([]*MapItem[int, string]{{1, "one"}, {2, "two"}})
	item, err := dict.PopItem()
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if item.Key != 2 || item.Value != "two" {
		t.Errorf("Expected popped item to be {2, 'two'}, got %v", item)
	}
	if len(dict.Keys()) != 1 {
		t.Errorf("Expected length of keys to be 1, got %d", len(dict.Keys()))
	}
}

func TestOrderedDict_PopFirst(t *testing.T) {
	dict := NewOrderedDict[int, string]([]*MapItem[int, string]{{1, "one"}, {2, "two"}})
	item, err := dict.PopFirst()
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if item.Key != 1 || item.Value != "one" {
		t.Errorf("Expected popped item to be {1, 'one'}, got %v", item)
	}
	if len(dict.Keys()) != 1 {
		t.Errorf("Expected length of keys to be 1, got %d", len(dict.Keys()))
	}
}

func TestOrderedDict_SetDefault(t *testing.T) {
	dict := NewOrderedDict[int, string]([]*MapItem[int, string]{{1, "one"}, {2, "two"}})
	val := dict.SetDefault(3, "three")
	if val != "three" {
		t.Errorf("Expected SetDefault to return 'three', got '%s'", val)
	}
	if len(dict.Keys()) != 3 {
		t.Errorf("Expected length of keys to be 3, got %d", len(dict.Keys()))
	}
}

func TestOrderedDict_Keys(t *testing.T) {
	dict := NewOrderedDict[int, string]([]*MapItem[int, string]{{1, "one"}, {2, "two"}})
	keys := dict.Keys()
	if len(keys) != 2 {
		t.Errorf("Expected length of keys to be 2, got %d", len(keys))
	}
	if !reflect.DeepEqual(keys, []int{1, 2}) {
		t.Errorf("Expected keys to be [1, 2], got %v", keys)
	}
}

func TestOrderedDict_SetItem(t *testing.T) {
	dict := NewOrderedDict[int, string]([]*MapItem[int, string]{{1, "one"}})
	dict.SetItem(2, "two")
	if len(dict.Keys()) != 2 {
		t.Errorf("Expected length of keys to be 2, got %d", len(dict.Keys()))
	}
	v, err := dict.GetItem(2)
	assert.Nil(t, err)
	if v != "two" {
		t.Errorf("Expected value for key 2 to be 'two', got '%s'", v)
	}
}

func TestOrderedDict_Delete(t *testing.T) {
	dict := NewOrderedDict[int, string]([]*MapItem[int, string]{{1, "one"}, {2, "two"}})
	dict.Delete(1)
	fmt.Printf("%v\n", dict.Keys())
	if len(dict.Keys()) != 1 {
		t.Errorf("Expected length of keys to be 1, got %d", len(dict.Keys()))
	}
	if dict.Contains(1) {
		t.Errorf("Expected key 1 to be deleted")
	}
}

func TestOrderedDict_Clear(t *testing.T) {
	dict := NewOrderedDict[int, string]([]*MapItem[int, string]{{1, "one"}, {2, "two"}})
	dict.Clear()
	if len(dict.Keys()) != 0 {
		t.Errorf("Expected length of keys to be 0, got %d", len(dict.Keys()))
	}
	if dict.Contains(1) || dict.Contains(2) {
		t.Errorf("Expected all keys to be deleted")
	}
}

func TestOrderedDict_FromKeys(t *testing.T) {
	dict := NewOrderedDict[int, string]()
	dict.FromKeys([]int{1, 2, 3}, "default")
	if len(dict.Keys()) != 3 {
		t.Errorf("Expected length of keys to be 3, got %d", len(dict.Keys()))
	}
	v, err := dict.GetItem(1)
	assert.Nil(t, err)
	if v != "default" {
		t.Errorf("Expected value for key 1 to be 'default', got '%s'", v)
	}
	v, err = dict.GetItem(3)
	assert.Nil(t, err)
	if v != "default" {
		t.Errorf("Expected value for key 3 to be 'default', got '%s'", v)
	}
}
