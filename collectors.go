package utils

import (
	"errors"
	"fmt"
	"strings"
)

type Set[T comparable] map[T]bool

func NewSet[T comparable](el ...T) Set[T] {
	newSet := make(Set[T])
	for _, v := range el {
		newSet.Add(v)
	}
	return newSet
}

func (s *Set[T]) FromSlice(t []T) {
	newSet := make(Set[T])
	for _, v := range t {
		newSet.Add(v)
	}
	*s = newSet
}

func (s Set[T]) ToSlice() []T {
	r := make([]T, 0, len(s))
	for v := range s {
		r = append(r, v)
	}
	return r
}

func (s Set[T]) String() string {
	r := Map(func(v T) string { return fmt.Sprintf("%v", v) }, s.ToSlice())
	return fmt.Sprintf("Set([ %s ])", strings.Join(r, " "))
}

func (s Set[T]) Contains(o T) bool {
	_, in := s[o]
	return in
}

func (s Set[T]) Add(o T) bool {
	if s.Contains(o) {
		return false
	}
	s[o] = true
	return true
}

func (s Set[T]) Discard(o T) bool {
	if !s.Contains(o) {
		return false
	}
	delete(s, o)
	return true
}

func (s Set[T]) Remove(o T) error {
	if !s.Contains(o) {
		return errors.New("element not found in set to remove")
	}
	delete(s, o)
	return nil
}

func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s Set[T]) Copy() Set[T] {
	newSet := make(Set[T])
	for k := range s {
		newSet.Add(k)
	}
	return newSet
}

func contains[T comparable](t T) func(Set[T]) bool {
	return func(s Set[T]) bool {
		return s.Contains(t)
	}
}

func (s Set[T]) Difference(sets ...Set[T]) Set[T] {
	newSet := make(Set[T])
	for k := range s {
		if !AnyFunc(contains(k), sets...) {
			newSet.Add(k)
		}
	}
	return newSet
}

func (s Set[T]) DifferenceUpdate(sets ...Set[T]) {
	for k := range s {
		if AnyFunc(contains(k), sets...) {
			s.Discard(k)
		}
	}
}

func (s Set[T]) Intersection(s1 Set[T]) Set[T] {
	newSet := make(Set[T])
	for o := range s {
		if s1.Contains(o) {
			newSet.Add(o)
		}
	}
	return newSet
}

func (s Set[T]) IntersectionSets(ss ...Set[T]) Set[T] {
	newSet := make(Set[T])
	for o := range s {
		if AllFunc(contains(o), ss...) {
			newSet.Add(o)
		}
	}
	return newSet
}

func (s Set[T]) IntersectionUpdate(s1 Set[T]) {
	for o := range s {
		if !s1.Contains(o) {
			s.Discard(o)
		}
	}
}

func (s Set[T]) IntersectionSetsUpdate(ss ...Set[T]) {
	for o := range s {
		if !AllFunc(contains(o), ss...) {
			s.Discard(o)
		}
	}
}

func (s Set[T]) IsDisjoint(s1 Set[T]) bool {
	for k := range s {
		if s1.Contains(k) {
			return false
		}
	}
	return true
}

// whether s1 contains all element of s
func (s Set[T]) IsSubset(s1 Set[T]) bool {
	for k := range s {
		if !s1.Contains(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) IsSuperset(s1 Set[T]) bool {
	return s1.IsSubset(s)
}

func (s Set[T]) Pop() (r T, err error) {
	for k := range s {
		s.Discard(k)
		return k, nil
	}
	return r, errors.New("pop from empty set")
}

func (s Set[T]) SymmetricDifference(s1 Set[T]) Set[T] {
	newSet := make(Set[T])
	for o := range s {
		if !s1.Contains(o) {
			newSet.Add(o)
		}
	}
	for o := range s1 {
		if !s.Contains(o) {
			newSet.Add(o)
		}
	}
	return newSet
}

func (s Set[T]) SymmetricDifferenceUpdate(s1 Set[T]) {
	inner := make(Set[T])
	for o := range s {
		if s1.Contains(o) {
			inner.Add(o)
			s.Remove(o)
		}
	}
	for o := range s1 {
		if !inner.Contains(o) {
			s.Add(o)
		}
	}
}

func (s Set[T]) Union(ss ...Set[T]) Set[T] {
	newSet := make(Set[T])
	for o := range s {
		newSet.Add(o)
	}
	for _, s2 := range ss {
		for o := range s2 {
			newSet.Add(o)
		}
	}
	return newSet
}

func (s Set[T]) Update(ss ...Set[T]) {
	for _, s2 := range ss {
		for o := range s2 {
			s.Add(o)
		}
	}
}
