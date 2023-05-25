package sort

import "sort"

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
