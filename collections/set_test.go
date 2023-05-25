package collections

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := NewSet(1, 2, 3)
	if len(s.m) != 3 {
		t.Errorf("NewSet should create a set with 3 elements")
	}
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("NewSet should create a set with elements {1, 2, 3}")
	}
}

func TestSetFromSlice(t *testing.T) {
	s := Set[int]{}
	s.FromSlice([]int{1, 2, 3})
	if len(s.m) != 3 {
		t.Errorf("FromSlice should create a set with 3 elements")
	}
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("FromSlice should create a set with elements {1, 2, 3}")
	}
}

func TestSetToSlice(t *testing.T) {
	s := NewSet(1, 2, 3)
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("ToSlice should return a slice with 3 elements")
	}
	if !assert.ElementsMatch(t, slice, []int{1, 2, 3}) {
		t.Errorf("ToSlice should return a slice with elements {1, 2, 3}")
	}
}

func TestSetString(t *testing.T) {
	s := NewSet(1, 2, 3)
	str := s.String()
	if str != "Set([ 1 2 3 ])" {
		t.Errorf("String should return 'Set([ 1 2 3 ])', but got %s", str)
	}
}

func TestSetContains(t *testing.T) {
	var s Set[int]
	s.FromSlice([]int{1, 2, 3})
	if !s.Contains(1) {
		t.Errorf("Set should contain element 1")
	}
	if s.Contains(4) {
		t.Errorf("Set should not contain element 4")
	}
}

func TestSetAdd(t *testing.T) {
	var s Set[int]
	s.FromSlice([]int{1, 2, 3})
	if s.Add(1) {
		t.Errorf("Set should not add existing element 1")
	}
	if !s.Add(4) {
		t.Errorf("Set should add new element 4")
	}
	if !s.Contains(4) {
		t.Errorf("Set should contain element 4")
	}
}

func TestSetDiscard(t *testing.T) {
	var s Set[int]
	s.FromSlice([]int{1, 2, 3})
	if !s.Discard(1) {
		t.Errorf("Set should discard existing element 1")
	}
	if s.Discard(4) {
		t.Errorf("Set should not discard non-existing element 4")
	}
	if s.Contains(1) {
		t.Errorf("Set should not contain discarded element 1")
	}
}

func TestSetRemove(t *testing.T) {
	var s Set[int]
	s.FromSlice([]int{1, 2, 3})
	if err := s.Remove(1); err != nil {
		t.Errorf("Set should remove existing element 1 without error")
	}
	if err := s.Remove(4); err == nil || !errors.Is(err, ErrElementNotFound) {
		t.Errorf("Set should return error for removing non-existing element 4")
	}
	if s.Contains(1) {
		t.Errorf("Set should not contain removed element 1")
	}
}

func TestSetClear(t *testing.T) {
	var s Set[int]
	s.FromSlice([]int{1, 2, 3})
	s.Clear()
	if len(s.m) != 0 {
		t.Errorf("Set should be empty after clearing")
	}
}

func TestSetCopy(t *testing.T) {
	var s Set[int]
	s.FromSlice([]int{1, 2, 3})
	copy := s.Copy()
	if !reflect.DeepEqual(s.m, copy.m) {
		t.Errorf("Copied set should be equal to original set")
	}
	copy.Add(4)
	if s.Contains(4) {
		t.Errorf("Original set should not be modified by copying")
	}
}

func TestSetDifference(t *testing.T) {
	var s1 Set[int]
	s1.FromSlice([]int{1, 2, 3})
	var s2 Set[int]
	s2.FromSlice([]int{2, 3, 4})
	diff := s1.Difference(&s2)
	expected := NewSet(1)
	if !reflect.DeepEqual(diff.m, expected.m) {
		t.Errorf("Difference should be %v, but got %v", expected, diff)
	}
}

func TestSetDifferenceUpdate(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)
	s1.DifferenceUpdate(s2)
	expected := NewSet(1)
	if !reflect.DeepEqual(s1.m, expected.m) {
		t.Errorf("DifferenceUpdate should modify set to %v, but got %v", expected, s1)
	}
}

func TestSetIntersection(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)
	inter := s1.Intersection(s2)
	expected := NewSet(2, 3)
	if !reflect.DeepEqual(inter.m, expected.m) {
		t.Errorf("Intersection should be %v, but got %v", expected, inter)
	}
}

func TestSetIntersectionSets(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)
	s3 := NewSet(3, 4, 5)
	inter := s1.IntersectionSets(s2, s3)
	expected := NewSet(3)
	if !reflect.DeepEqual(inter.m, expected.m) {
		t.Errorf("IntersectionSets should be %v, but got %v", expected, inter)
	}
}

func TestSetIntersectionUpdate(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)
	s1.IntersectionUpdate(s2)
	expected := NewSet(2, 3)
	if !reflect.DeepEqual(s1.m, expected.m) {
		t.Errorf("IntersectionUpdate should modify set to %v, but got %v", expected, s1)
	}
}

func TestSetIntersectionSetsUpdate(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)
	s3 := NewSet(3, 4, 5)
	s1.IntersectionSetsUpdate(s2, s3)
	expected := NewSet(3)
	if !reflect.DeepEqual(s1.m, expected.m) {
		t.Errorf("IntersectionSetsUpdate should modify set to %v, but got %v", expected, s1)
	}
}

func TestSetIsDisjoint(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(4, 5, 6)
	if !s1.IsDisjoint(s2) {
		t.Errorf("Sets should be disjoint")
	}
	s3 := NewSet(3, 4, 5)
	if s1.IsDisjoint(s3) {
		t.Errorf("Sets should not be disjoint")
	}
}

func TestSetIsSubset(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(1, 2, 3, 4, 5)
	if !s1.IsSubset(s2) {
		t.Errorf("Set should be subset")
	}
	s3 := NewSet(1, 2)
	if s1.IsSubset(s3) {
		t.Errorf("Set should not be subset")
	}
}

func TestSetIsSuperset(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(1, 2, 3, 4, 5)
	if !s2.IsSuperset(s1) {
		t.Errorf("Set should be superset")
	}
	s3 := NewSet(1, 2)
	if s3.IsSuperset(s1) {
		t.Errorf("Set should not be superset")
	}
}

func TestSetPop(t *testing.T) {
	s := NewSet(1, 2, 3)
	sCopy := s.Copy()
	p, err := s.Pop()
	if err != nil || s.Contains(p) {
		t.Errorf("Pop should remove the popped element")
	}
	if !sCopy.Contains(p) {
		t.Errorf("Pop should return an element from set")
	}
	for i := 0; i < 2; i++ {
		_, _ = s.Pop()
	}
	_, err = s.Pop()
	if err == nil || !errors.Is(err, ErrPopFromEmptyCollections) {
		t.Errorf("Pop should return error for empty set")
	}
}

func TestSetSymmetricDifference(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)
	symDiff := s1.SymmetricDifference(s2)
	expected := NewSet(1, 4)
	if !reflect.DeepEqual(symDiff, expected) {
		t.Errorf("SymmetricDifference should be %v, but got %v", expected, symDiff)
	}
}

func TestSetSymmetricDifferenceUpdate(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)
	s1.SymmetricDifferenceUpdate(s2)
	expected := NewSet(1, 4)
	if !reflect.DeepEqual(s1, expected) {
		t.Errorf("SymmetricDifferenceUpdate should modify set to %v, but got %v", expected, s1)
	}
}

func TestSetUnion(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)
	union := s1.Union(s2)
	expected := NewSet(1, 4, 2, 3)
	if !reflect.DeepEqual(union.m, expected.m) {
		t.Errorf("Union should be %v, but got %v", expected, union)
	}
}

func TestSetUpdate(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)
	s1.Update(s2)
	expected := NewSet(1, 4, 2, 3)
	if !reflect.DeepEqual(s1, expected) {
		t.Errorf("Update should modify set to %v, but got %v", expected, s1)
	}
}
