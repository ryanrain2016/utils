package sort

import (
	"fmt"
	"sort"

	"github.com/ryanrain2016/utils/functools"
)

type ISort[T comparable] interface {
	Len() int
	Swap(i, j int)
	GetItem(i int) (T, error)
}

type lessFunc[T comparable] func(T, T) bool

type Sorter[T comparable] struct {
	less  []lessFunc[T]
	items ISort[T]
}

func SortBy[T comparable](less ...lessFunc[T]) *Sorter[T] {
	if len(less) == 0 {
		panic("should has at least one less function")
	}
	return &Sorter[T]{
		less: less,
	}
}

func SortByKey[T comparable](key ...any) *Sorter[T] {
	if len(key) == 0 {
		panic("should has at least one key function")
	}
	less := make([]lessFunc[T], len(key))
	for i, kf := range key {
		switch f := kf.(type) {
		case func(T) int:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) float32:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) float64:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) int8:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) int16:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) int32:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) int64:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) uint:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) uint8:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) uint16:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) uint32:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) uint64:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		case func(T) string:
			less[i] = func(t1, t2 T) bool {
				return f(t1) < f(t2)
			}
		default:
			sig := functools.FuncSignature(kf)
			panic(fmt.Sprintf("unsupport key function: %s", sig))
		}
	}
	return &Sorter[T]{
		less: less,
	}
}

func (s *Sorter[T]) Len() int {
	return s.items.Len()
}

func (s *Sorter[T]) Swap(i, j int) {
	s.items.Swap(i, j)
}

func (s *Sorter[T]) Less(i, j int) bool {
	p, _ := s.items.GetItem(i)
	q, _ := s.items.GetItem(j)
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(s.less)-1; k++ {
		less := s.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return s.less[k](p, q)
}

func (s *Sorter[T]) Sort(items ISort[T]) {
	s.items = items
	sort.Sort(s)
}
