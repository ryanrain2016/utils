package collections

import (
	"reflect"
	"testing"
)

func TestMapItem(t *testing.T) {
	item1 := &MapItem[int, string]{Key: 1, Value: "one"}
	item2 := &MapItem[int, string]{Key: 2, Value: "two"}

	// Test equality and inequality of MapItem instances
	if !reflect.DeepEqual(item1, item1) {
		t.Errorf("Expected item1 to be equal to itself")
	}
	if reflect.DeepEqual(item1, item2) {
		t.Errorf("Expected item1 and item2 to be unequal")
	}
}

func TestNewDict(t *testing.T) {
	// Test empty dictionary
	dict := NewDict[int, string]()
	if len(dict.Map()) != 0 {
		t.Errorf("Expected empty dictionary")
	}

	// Test dictionary with multiple key-value pairs
	dict = NewDict[int, string](map[int]string{1: "one", 2: "two"})
	if len(dict.Map()) != 2 {
		t.Errorf("Expected dictionary length to be 2")
	}
	if dict.Get(1) != "one" || dict.Get(2) != "two" {
		t.Errorf("Expected dictionary values to be 'one' and 'two'")
	}

	// Test dictionary with duplicate keys
	dict = NewDict[int, string](map[int]string{1: "one"})
	dict.SetItem(1, "uno")
	if len(dict.Map()) != 1 {
		t.Errorf("Expected dictionary length to be 1")
	}
	if dict.Get(1) != "uno" {
		t.Errorf("Expected dictionary value to be 'uno'")
	}
}

func TestMapMethod(t *testing.T) {
	dict := NewDict[int, string](map[int]string{1: "one", 2: "two"})
	m := dict.Map()
	if len(m) != 2 {
		t.Errorf("Expected map length to be 2")
	}
	if m[1] != "one" || m[2] != "two" {
		t.Errorf("Expected map values to be 'one' and 'two'")
	}
}

func TestCopyMethod(t *testing.T) {
	dict := NewDict[int, string](map[int]string{1: "one", 2: "two"})
	copy := dict.Copy()
	if !reflect.DeepEqual(dict.Map(), copy.Map()) {
		t.Errorf("Expected copy to have the same contents as the original")
	}
	copy.SetItem(3, "three")
	if dict.Get(3) == "three" {
		t.Errorf("Expected copy to be a different object than the original")
	}
}

func TestDict(t *testing.T) {
	// test FromKeys method
	d := NewDict[int, string]()
	keys := []int{1, 2, 3}
	d.FromKeys(keys, "a")
	if d.Len() != 3 {
		t.Errorf("Expected length to be %d, got %d", 3, d.Len())
	}
	if d.Get(1) != "a" || d.Get(2) != "a" || d.Get(3) != "a" {
		t.Errorf("Expected values to be %v, got %v", []string{"a", "a", "a"}, []string{d.Get(1), d.Get(2), d.Get(3)})
	}

	// test Get method
	if d.Get(4) != "" {
		t.Errorf("Expected value to be empty string, got %s", d.Get(4))
	}
	if d.Get(4, "d") != "d" {
		t.Errorf("Expected value to be %s, got %s", "d", d.Get(4, "d"))
	}

	// test Items method
	items := d.Items()
	if len(items) != 3 {
		t.Errorf("Expected length to be %d, got %d", 3, len(items))
	}
	// test Keys method
	keys = d.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected length to be %d, got %d", 3, len(keys))
	}
	if !sliceContains(keys, 1) || !sliceContains(keys, 2) || !sliceContains(keys, 3) {
		t.Errorf("Expected keys to contain %v, got %v", []int{1, 2, 3}, keys)
	}

	// test Pop method
	v, err := d.Pop(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if v != "a" {
		t.Errorf("Expected value to be %s, got %s", "a", v)
	}
	if d.Len() != 2 {
		t.Errorf("Expected length to be %d, got %d", 2, d.Len())
	}
	if d.Get(1) != "" {
		t.Errorf("Expected value to be empty string, got %s", d.Get(1))
	}
	v, err = d.Pop(4)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if v != "" {
		t.Errorf("Expected value to be empty string, got %s", v)
	}

	// test PopItem method
	cp := d.Copy()
	item, err := d.PopItem()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !cp.Contains(item.Key) || item.Value != cp.Get(item.Key) {
		t.Errorf("Expected item to be {%d, %s}, got {%d, %s}", item.Key, cp.Get(item.Key), item.Key, item.Value)
	}
	if d.Len() != 1 {
		t.Errorf("Expected length to be %d, got %d", 1, d.Len())
	}

	// test SetDefault method
	d4 := NewDict[int, string](map[int]string{1: "a", 2: "b"})
	if d4.SetDefault(2) != "b" {
		t.Errorf("Expected value to be %s, got %s", "b", d.SetDefault(2))
	}
	if v := d4.SetDefault(4, "d"); v != "d" {
		t.Errorf("Expected value to be %s, got %s", "d", v)
	}
	if d4.Get(4) != "d" {
		t.Errorf("Expected value to be %s, got %s", "d", d.Get(4))
	}

	// test Update Items
	d = NewDict[int, string](map[int]string{1: "a", 4: "b"})
	items = []*MapItem[int, string]{{Key: 1, Value: "f"}, {Key: 4, Value: "g"}}
	d.Update(items)
	if d.Len() != 2 {
		t.Errorf("Expected length to be %d, got %d", 2, d.Len())
	}
	if d.Get(1) != "f" || d.Get(4) != "g" {
		t.Errorf("Expected values to be %v, got %v", []string{"f", "g"}, []string{d.Get(1), d.Get(4)})
	}

	// test UpdatePairSlice method
	d = NewDict[int, string](map[int]string{1: "a", 4: "b"})
	keys = []int{1, 5}
	values := []string{"h", "i"}
	d.UpdatePairSlice(keys, values)
	if d.Len() != 3 {
		t.Errorf("Expected length to be %d, got %d", 3, d.Len())
	}
	if d.Get(1) != "h" || d.Get(5) != "i" {
		t.Errorf("Expected values to be %v, got %v", []string{"h", "i"}, []string{d.Get(1), d.Get(5)})
	}

	// test Update Map
	d = NewDict[int, string](map[int]string{2: "a", 4: "b"})
	m := map[int]string{1: "j", 6: "k"}
	d.Update(m)
	if d.Len() != 4 {
		t.Errorf("Expected length to be %d, got %d", 4, d.Len())
	}
	if d.Get(1) != "j" || d.Get(6) != "k" {
		t.Errorf("Expected values to be %v, got %v", []string{"j", "k"}, []string{d.Get(1), d.Get(6)})
	}

	// test Update method
	d = NewDict[int, string](map[int]string{1: "a", 3: "c", 4: "b"})
	d2 := NewDict[int, string]()
	d2.SetItem(2, "l")
	d.Update(d2)
	if d.Len() != 4 {
		t.Errorf("Expected length to be %d, got %d", 4, d.Len())
	}
	if d.Get(2) != "l" {
		t.Errorf("Expected value to be %s, got %s", "l", d.Get(2))
	}

	// test Len method
	d3 := NewDict[int, string]()
	d3.SetItem(1, "abc")
	d3.SetItem(2, "cd")
	if d3.Len() != 2 {
		t.Errorf("Expected value to be %d, got %d", 2, d3.Len())
	}

	// test Set method
	d.SetItem(1, "m")
	if d.Get(1) != "m" {
		t.Errorf("Expected value to be %s, got %s", "m", d.Get(1))
	}

	// test Delete method
	d = NewDict[int, string](map[int]string{1: "a", 3: "c", 4: "b", 5: "h"})
	d.Delete(5)
	if d.Len() != 3 {
		t.Errorf("Expected length to be %d, got %d", 3, d.Len())
	}
	if d.Get(5) != "" {
		t.Errorf("Expected value to be empty string, got %s", d.Get(5))
	}
}

func sliceContains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func TestDictEq(t *testing.T) {
	d := NewDict[int, string](map[int]string{1: "a", 3: "c", 4: "b", 5: "h"})
	d1 := NewDict[int, string](map[int]string{1: "a", 3: "c", 4: "b", 5: "h"})
	if !d.Eq(d1) {
		t.Errorf("the to Dict are Expected to be equal")
	}
}
