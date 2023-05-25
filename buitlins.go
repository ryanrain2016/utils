package utils

import (
	"errors"
	"runtime"
	"sync"

	"github.com/ryanrain2016/utils/types"
)

func Map[T any, U any](mapFunc func(T) U, slice []T) []U {
	rslt := make([]U, 0)
	for _, v := range slice {
		rslt = append(rslt, mapFunc(v))
	}
	return rslt
}

func MapToChan[T any, U any](mapFunc func(T) U, slice []T) chan U {
	rslt := make(chan U)
	go func() {
		defer func() {
			close(rslt)
		}()
		for _, v := range slice {
			rslt <- mapFunc(v)
		}
	}()
	return rslt
}

func ConcurrentMap[T any, U any](mapFunc func(T) U, slice []T) []U {
	n := len(slice)
	maxWorkers := runtime.NumCPU()
	if n < maxWorkers {
		return Map(mapFunc, slice)
	}
	chunkSize := (n + maxWorkers - 1) / maxWorkers
	rslt := make([]U, n)
	var wg sync.WaitGroup
	wg.Add(maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > n {
			end = n
		}

		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				rslt[j] = mapFunc(slice[j])
			}
		}(start, end)
	}
	wg.Wait()
	return rslt
}

func Max[T types.Ordered](items ...T) (*T, error) {
	if len(items) == 0 {
		return nil, errors.New("Max should be called with at least one arguments")
	}
	max := &items[0]
	for i := 1; i < len(items); i++ {
		if items[i] > *max {
			max = &items[i]
		}
	}
	return max, nil
}

func MaxSlice[T types.Ordered](items []T) (*T, error) {
	if len(items) == 0 {
		return nil, errors.New("MaxSlice should be called with slice contains at least one item")
	}
	return Max(items...)
}

func Min[T types.Ordered](items ...T) (*T, error) {
	if len(items) == 0 {
		return nil, errors.New("Min should be called with at least one arguments")
	}
	min := &items[0]
	for i := 1; i < len(items); i++ {
		if items[i] < *min {
			min = &items[i]
		}
	}
	return min, nil
}

func MinSlice[T types.Ordered](items []T) (*T, error) {
	if len(items) == 0 {
		return nil, errors.New("MinSlice should be called with slice contains at least one item")
	}
	return Min(items...)
}

func Any(items ...bool) bool {
	for _, v := range items {
		if v {
			return true
		}
	}
	return false
}

func AnyFunc[T any](f func(T) bool, items ...T) bool {
	for _, v := range items {
		if f(v) {
			return true
		}
	}
	return false
}

func AnyFuncSlice[T any](f func(T) bool, items []T) bool {
	return AnyFunc(f, items...)
}

func All(items ...bool) bool {
	for _, v := range items {
		if !v {
			return false
		}
	}
	return true
}

func AllFunc[T any](f func(T) bool, items ...T) bool {
	for _, v := range items {
		if !f(v) {
			return false
		}
	}
	return true
}

func AllFuncSlice[T any](f func(T) bool, items []T) bool {
	return AllFunc(f, items...)
}

func Sum[T types.Addable](items ...T) T {
	var r T
	for _, v := range items {
		r += v
	}
	return r
}

func SumSlice[T types.Addable](items []T) T {
	return Sum(items...)
}
