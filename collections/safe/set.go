package safe

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ryanrain2016/utils"
	"github.com/ryanrain2016/utils/collections"
)

type Set[T comparable] struct {
	m map[T]bool
	sync.RWMutex
}

func NewSet[T comparable](el ...T) *Set[T] {
	newSet := new(Set[T])
	newSet.m = make(map[T]bool)
	for _, v := range el {
		newSet.Add(v)
	}
	return newSet
}

func (s *Set[T]) FromSlice(t []T) {
	newSet := NewSet(t...)
	s.m = newSet.m
}

func (s *Set[T]) ToSlice() []T {
	s.RLock()
	defer s.RUnlock()
	r := make([]T, 0, len(s.m))
	for v := range s.m {
		r = append(r, v)
	}
	return r
}

func (s *Set[T]) String() string {
	r := utils.Map(func(v T) string { return fmt.Sprintf("%v", v) }, s.ToSlice())
	return fmt.Sprintf("Set([ %s ])", strings.Join(r, " "))
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Contains(o T) bool {
	s.RLock()
	defer s.RUnlock()
	_, in := s.m[o]
	return in
}

func (s *Set[T]) Add(o T) bool {
	if s.Contains(o) {
		return false
	}
	s.Lock()
	defer s.Unlock()
	s.m[o] = true
	return true
}

func (s *Set[T]) Discard(o T) bool {
	if !s.Contains(o) {
		return false
	}
	s.Lock()
	defer s.Unlock()
	delete(s.m, o)
	return true
}

func (s *Set[T]) Remove(o T) error {
	if !s.Contains(o) {
		return collections.ErrElementNotFound
	}
	s.Lock()
	defer s.Unlock()
	delete(s.m, o)
	return nil
}

func (s *Set[T]) Clear() {
	s.Lock()
	defer s.Unlock()
	for k := range s.m {
		delete(s.m, k)
	}
}

func (s *Set[T]) Copy() *Set[T] {
	s.RLock()
	defer s.RUnlock()
	newSet := NewSet[T]()
	for k := range s.m {
		newSet.Add(k)
	}
	return newSet
}

func contains[T comparable](t T) func(*Set[T]) bool {
	return func(s *Set[T]) bool {
		return s.Contains(t)
	}
}

func (s *Set[T]) Difference(sets ...*Set[T]) *Set[T] {
	newSet := NewSet[T]()
	s.RLock()
	defer s.RUnlock()
	for k := range s.m {
		if !utils.AnyFunc(contains(k), sets...) {
			newSet.m[k] = true
		}
	}
	return newSet
}

func (s *Set[T]) DifferenceUpdate(sets ...*Set[T]) {
	s.Lock()
	defer s.Unlock()
	for k := range s.m {
		if utils.AnyFunc(contains(k), sets...) {
			delete(s.m, k)
		}
	}
}

func (s *Set[T]) Intersection(s1 *Set[T]) *Set[T] {
	newSet := NewSet[T]()
	s.Lock()
	defer s.Unlock()
	for o := range s.m {
		if s1.Contains(o) {
			newSet.Add(o)
		}
	}
	return newSet
}

func (s *Set[T]) IntersectionSets(ss ...*Set[T]) *Set[T] {
	newSet := NewSet[T]()
	s.RLock()
	defer s.RUnlock()
	for o := range s.m {
		if utils.AllFunc(contains(o), ss...) {
			newSet.Add(o)
		}
	}
	return newSet
}

func (s *Set[T]) IntersectionUpdate(s1 *Set[T]) {
	s.Lock()
	defer s.Unlock()
	for o := range s.m {
		if !s1.Contains(o) {
			delete(s.m, o)
		}
	}
}

func (s *Set[T]) IntersectionSetsUpdate(ss ...*Set[T]) {
	for o := range s.m {
		if !utils.AllFunc(contains(o), ss...) {
			delete(s.m, o)
		}
	}
}

func (s *Set[T]) IsDisjoint(s1 *Set[T]) bool {
	s.RLock()
	defer s.RUnlock()
	for k := range s.m {
		if s1.Contains(k) {
			return false
		}
	}
	return true
}

// whether s1 contains all element of s
func (s *Set[T]) IsSubset(s1 *Set[T]) bool {
	s.RLock()
	defer s.RUnlock()
	for k := range s.m {
		if !s1.Contains(k) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSuperset(s1 *Set[T]) bool {
	return s1.IsSubset(s)
}

func (s *Set[T]) Pop() (r T, err error) {
	s.Lock()
	defer s.Unlock()
	for k := range s.m {
		delete(s.m, k)
		return k, nil
	}
	return r, collections.ErrPopFromEmptyCollections
}

func (s *Set[T]) SymmetricDifference(s1 *Set[T]) *Set[T] {
	newSet := NewSet[T]()
	s.RLock()
	defer s.RUnlock()
	for o := range s.m {
		if !s1.Contains(o) {
			newSet.Add(o)
		}
	}
	s1.Lock()
	defer s1.Unlock()
	for o := range s1.m {
		if !s.Contains(o) {
			newSet.Add(o)
		}
	}
	return newSet
}

func (s *Set[T]) SymmetricDifferenceUpdate(s1 *Set[T]) {
	inner := make(map[T]bool)
	s.RLock()
	defer s.RUnlock()
	for o := range s.m {
		if s1.Contains(o) {
			inner[o] = true
			delete(s.m, o)
		}
	}
	s1.RLock()
	defer s1.RUnlock()
	for o := range s1.m {
		if !inner[o] {
			s.m[o] = true
		}
	}
}

func (s *Set[T]) Union(ss ...*Set[T]) *Set[T] {
	newSet := NewSet[T]()
	s.RLock()
	defer s.RUnlock()
	for o := range s.m {
		newSet.m[o] = true
	}
	wg := &sync.WaitGroup{}
	for _, s2 := range ss {
		wg.Add(1)
		go func(s2 *Set[T]) {
			defer wg.Done()
			s2.RLock()
			defer s2.RUnlock()
			for o := range s2.m {
				newSet.Add(o)
			}
		}(s2)
	}
	wg.Wait()
	return newSet
}

func (s *Set[T]) Update(ss ...*Set[T]) {
	wg := &sync.WaitGroup{}
	for _, s2 := range ss {
		wg.Add(1)
		go func(s2 *Set[T]) {
			defer wg.Done()
			s2.RLock()
			defer s2.RUnlock()
			for o := range s2.m {
				s.Lock()
				s.m[o] = true
				s.Unlock()
			}
		}(s2)
	}
	wg.Wait()
}
