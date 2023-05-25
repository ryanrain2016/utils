package collections

import "testing"

func TestNewCounter(t *testing.T) {
	counter := NewCounter[int]()
	if counter == nil {
		t.Errorf("NewCounter() should not return nil")
		return
	}
	if counter.m == nil || len(counter.m) != 0 {
		t.Errorf("NewCounter() should initialize an empty map")
	}
	if counter.items == nil || len(counter.items) != 0 {
		t.Errorf("NewCounter() should initialize an empty slice")
	}
}

func TestCounterFromSlice(t *testing.T) {
	counter := NewCounter[int]()
	counter.FromSlice([]int{1, 1, 2, 3})
	if len(counter.m) != 3 {
		t.Errorf("FromSlice() should add all elements to the map")
	}
	if len(counter.items) != 3 {
		t.Errorf("FromSlice() should add unique elements to the slice")
	}
	if counter.m[1] != 2 || counter.m[2] != 1 || counter.m[3] != 1 {
		t.Errorf("FromSlice() should count the frequency of each element")
	}
}

func TestCounterFromMap(t *testing.T) {
	counter := NewCounter[int]()
	counter.FromMap(map[int]int{1: 2, 2: 1, 3: 1})
	if len(counter.m) != 3 {
		t.Errorf("FromMap() should add all elements to the map")
	}
	if len(counter.items) != 3 {
		t.Errorf("FromMap() should add unique elements to the slice")
	}
	if counter.m[1] != 2 || counter.m[2] != 1 || counter.m[3] != 1 {
		t.Errorf("FromMap() should set the count of each element")
	}
}

func TestCounterCopy(t *testing.T) {
	counter1 := NewCounter[int]()
	counter1.Count(1)
	counter1.Count(2)
	counter2 := counter1.Copy().(*Counter[int])
	if len(counter2.m) != 2 {
		t.Errorf("Copy() should copy the map")
	}
	if len(counter2.items) != 2 {
		t.Errorf("Copy() should copy the slice")
	}
	counter1.Count(3)
	if len(counter2.m) != 2 {
		t.Errorf("Copy() should create a new map")
	}
	if len(counter2.items) != 2 {
		t.Errorf("Copy() should create a new slice")
	}
}

func TestCounterCount(t *testing.T) {
	counter := NewCounter[string]()
	counter.Count("a")
	if len(counter.m) != 1 {
		t.Errorf("Count() should add new elements to the map")
	}
	if len(counter.items) != 1 {
		t.Errorf("Count() should add new elements to the slice")
	}
	if counter.m["a"] != 1 {
		t.Errorf("Count() should increase the count of an existing element")
	}
	counter.Count("a")
	if counter.m["a"] != 2 {
		t.Errorf("Count() should increase the count of an existing element")
	}
}

func TestCounterElements(t *testing.T) {
	counter := NewCounter[int]()
	counter.Count(1)
	counter.Count(1)
	counter.Count(2)
	counter.Count(3)
	elements := counter.Elements()
	if len(elements) != 4 {
		t.Errorf("Elements() should return all elements")
	}
	if elements[0] != 1 || elements[1] != 1 || elements[2] != 2 || elements[3] != 3 {
		t.Errorf("Elements() should return elements in the order they were counted")
	}
}

func TestCounterMostCommon(t *testing.T) {
	counter := NewCounter[string]()
	counter.Count("a")
	counter.Count("b")
	counter.Count("c")
	counter.Count("a")
	mostCommon := counter.MostCommon()
	if len(mostCommon) != 3 {
		t.Errorf("MostCommon() should return all elements")
	}
	if mostCommon[0].Key != "a" || mostCommon[0].Value != 2 {
		t.Errorf("MostCommon() should return elements sorted by count")
	}
	if mostCommon[1].Key != "b" || mostCommon[1].Value != 1 {
		t.Errorf("MostCommon() should return elements sorted by count")
	}
	if mostCommon[2].Key != "c" || mostCommon[2].Value != 1 {
		t.Errorf("MostCommon() should return elements sorted by count")
	}
	mostCommon = counter.MostCommon(2)
	if len(mostCommon) != 2 {
		t.Errorf("MostCommon() should return specified number of elements")
	}
	if mostCommon[0].Key != "a" || mostCommon[0].Value != 2 {
		t.Errorf("MostCommon() should return specified number of elements sorted by count")
	}
	if mostCommon[1].Key != "b" || mostCommon[1].Value != 1 {
		t.Errorf("MostCommon() should return specified number of elements sorted by count")
	}
}

func TestCounterSubtract(t *testing.T) {
	counter1 := NewCounter[int]()
	counter1.Count(1)
	counter1.Count(1)
	counter1.Count(2)
	counter1.Count(3)
	counter2 := NewCounter[int]()
	counter2.Count(1)
	counter2.Count(2)
	counter1.Subtract(counter2)
	if counter1.m[1] != 1 || counter1.m[2] != 0 || counter1.m[3] != 1 {
		t.Errorf("Subtract() should subtract the counts of common elements")
	}
}

func TestCounterTotal(t *testing.T) {
	counter := NewCounter[int]()
	counter.Count(1)
	counter.Count(1)
	counter.Count(2)
	counter.Count(3)
	total := counter.Total()
	if total != 4 {
		t.Errorf("Total() should return the total count")
	}
}

func TestCounterUpdate(t *testing.T) {
	// Test updating with a slice
	counter1 := NewCounter[int]()
	counter1.Count(1)
	counter1.Count(2)
	counter2 := []int{1, 1, 3}
	counter1.Update(counter2)
	if counter1.Get(1) != 3 || counter1.Get(2) != 1 || counter1.Get(3) != 1 {
		t.Errorf("Update() should correctly update the counter with a slice")
	}

	// Test updating with a Counter
	counter1 = NewCounter[int]()
	counter1.Count(1)
	counter1.Count(2)
	counter3 := NewCounter[int]()
	counter3.Count(1)
	counter3.Count(1)
	counter3.Count(3)
	counter1.Update(counter3)
	if counter1.Get(1) != 3 || counter1.Get(2) != 1 || counter1.Get(3) != 1 {
		t.Errorf("Update() should correctly update the counter with a Counter")
	}

	// Test updating with a map
	counter1 = NewCounter[int]()
	counter1.Count(1)
	counter1.Count(2)
	counter4 := map[int]int{1: 2, 2: 1, 3: 3}
	counter1.Update(counter4)
	if counter1.Get(1) != 3 || counter1.Get(2) != 2 || counter1.Get(3) != 3 {
		t.Errorf("Update() should correctly update the counter with a map")
	}
}

func TestCounterToMap(t *testing.T) {
	counter := NewCounter[int]()
	counter.Count(1)
	counter.Count(1)
	counter.Count(2)
	otherMap := counter.toMap(counter)
	if otherMap[1] != 2 || otherMap[2] != 1 {
		t.Errorf("toMap() should return the map of the other Counter")
	}
	otherCounter := NewCounter[int]()
	otherCounter.Count(1)
	otherMap = counter.toMap(otherCounter)
	if otherMap[1] != 1 {
		t.Errorf("toMap() should return the map of the other Counter")
	}
	otherMap = counter.toMap(map[int]int{1: 2, 2: 1})
	if otherMap[1] != 2 || otherMap[2] != 1 {
		t.Errorf("toMap() should return the given map")
	}
	otherMap = counter.toMap([]int{1, 2})
	if otherMap[1] != 1 || otherMap[2] != 1 {
		t.Errorf("toMap() should count the frequency of elements in the slice")
	}
}

func TestCounterCounts(t *testing.T) {
	// Test counting a single element
	counter := NewCounter[int]()
	counter.Count(1)
	counter.Count(1)
	counter.Count(2)
	if counter.Get(1) != 2 {
		t.Errorf("Counts() should correctly count a single element")
	}
	if counter.Get(2) != 1 {
		t.Errorf("Counts() should correctly count a single element")
	}
	if counter.Get(3) != 0 {
		t.Errorf("Counts() should return 0 for an element not in the counter")
	}
}
