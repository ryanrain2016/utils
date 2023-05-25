package itertools

import (
	"reflect"

	"github.com/ryanrain2016/utils"
)

func Zip[T any](items ...[]T) (r [][]T) {
	minLength, err := utils.MinSlice(utils.Map(func(item []T) int { return len(item) }, items))
	if err != nil {
		return r
	}
	r = make([][]T, *minLength)
	for i := 0; i < *minLength; i++ {
		item := make([]T, len(items))
		for j, v := range items {
			item[j] = v[i]
		}
		r[i] = item
	}
	return r
}
func ZipSlice[T any](items [][]T) [][]T {
	return Zip(items...)
}

func ZipToChan[T any](items ...[]T) chan []T {
	r := make(chan []T)
	go func() {
		defer func() {
			close(r)
		}()
		minLength, err := utils.MinSlice(utils.Map(func(item []T) int { return len(item) }, items))
		if err != nil {
			return
		}
		for i := 0; i < *minLength; i++ {
			item := make([]T, len(items))
			for j, v := range items {
				item[j] = v[i]
			}
			r <- item
		}
	}()
	return r
}

func ZipSliceToChan[T any](items [][]T) chan []T {
	return ZipToChan(items...)
}

func ZipAny(items ...[]any) (r [][]any) {
	minLength, err := utils.MinSlice(utils.Map(func(item []any) int { return len(item) }, items))
	if err != nil {
		return r
	}
	r = make([][]any, *minLength)
	for i := 0; i < *minLength; i++ {
		item := make([]any, len(items))
		for j, v := range items {
			item[j] = v[i]
		}
		r[i] = item
	}
	return r
}

func ZipLongest[T any](items ...[]T) (r [][]T) {
	maxLength, err := utils.MaxSlice(utils.Map(func(item []T) int { return len(item) }, items))
	if err != nil {
		return r
	}
	r = make([][]T, *maxLength)
	for i := 0; i < *maxLength; i++ {
		item := make([]T, len(items))
		for j, v := range items {
			if i < len(v) {
				item[j] = v[i]
			}
		}
		r[i] = item
	}
	return r
}

func ZipLongestSlice[T any](items [][]T) [][]T {
	return ZipLongest(items...)
}

func ZipLongestToChan[T any](items ...[]T) chan []T {
	r := make(chan []T)
	go func() {
		defer func() {
			close(r)
		}()
		maxLength, err := utils.MaxSlice(utils.Map(func(item []T) int { return len(item) }, items))
		if err != nil {
			return
		}
		for i := 0; i < *maxLength; i++ {
			item := make([]T, len(items))
			for j, v := range items {
				if i < len(v) {
					item[j] = v[i]
				}
			}
			r <- item
		}
	}()
	return r
}

func ZipLongestSliceToChan[T any](items [][]T) chan []T {
	return ZipLongestToChan(items...)
}

func Filter[T any](filter func(T) bool, slice []T) []T {
	r := make([]T, 0, len(slice))
	for _, v := range slice {
		if filter(v) {
			r = append(r, v)
		}
	}
	return r
}

func FilterToChan[T any](filter func(T) bool, slice []T) chan T {
	r := make(chan T)
	go func() {
		defer close(r)
		for _, v := range slice {
			if filter(v) {
				r <- v
			}
		}
	}()
	return r
}

func FindFunc[T any](find func(T) bool, slice []T) (r T) {
	for _, v := range slice {
		if find(v) {
			return v
		}
	}
	return r
}

func FindIndexFunc[T any](find func(T) bool, slice []T) int {
	for i, v := range slice {
		if find(v) {
			return i
		}
	}
	return -1
}

func Accumulate[T any](slice []T, f func(T, T) T) chan T {
	r := make(chan T)
	go func() {
		defer close(r)
		if len(slice) == 0 {
			return
		}
		result := slice[0]
		r <- result
		for _, v := range slice[1:] {
			result = f(result, v)
			r <- result
		}
	}()
	return r
}

func Chain[T any](iters ...[]T) chan T {
	r := make(chan T)
	go func() {
		defer close(r)
		for _, iter := range iters {
			for _, v := range iter {
				r <- v
			}
		}
	}()
	return r
}

func FromIterable[T any](iterable [][]T) chan T {
	return Chain(iterable...)
}

func Combinations[T any](seq []T, r int) chan []T {
	if r <= 0 {
		panic("r must be greater than zero")
	}
	result := make(chan []T)
	if r > len(seq) {
		defer close(result)
		return result
	}
	go func() {
		defer close(result)
		indices := make([]int, r)
		for i := range indices {
			indices[i] = i
		}
		for {
			comb := make([]T, r)
			for i, j := range indices {
				comb[i] = seq[j]
			}
			result <- comb

			// Find the first index k from the right where indices[k] != len(seq) - r + k
			k := r - 1
			for ; k >= 0 && indices[k] == len(seq)-r+k; k-- {
			}

			if k < 0 {
				break
			}

			indices[k] += 1
			for i := k + 1; i < r; i++ {
				indices[i] = indices[i-1] + 1
			}
		}
	}()
	return result
}

func FindIndex[T comparable](seq []T, elem T) int {
	for i, e := range seq {
		if reflect.DeepEqual(e, elem) {
			return i
		}
	}
	return -1
}

func CombinationsWithReplacement[T comparable](seq []T, r int) chan []T {
	if r <= 0 {
		panic("r must be greater than zero")
	}
	result := make(chan []T)
	go func() {
		defer close(result)
		n := len(seq)
		indices := make([]int, r)
		for {
			comb := make([]T, r)
			for i, j := range indices {
				comb[i] = seq[j]
			}
			result <- comb

			// Increment the right-most index that is not equal to n-1
			k := r - 1
			for ; k >= 0 && indices[k] == n-1; k-- {
			}

			if k < 0 {
				break
			}
			index := indices[k] + 1
			for i := k; i < r; i++ {
				indices[i] = index
			}
		}
	}()

	return result
}

func Compress[T any](seq []T, selectors []bool) chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		for i, s := range seq {
			if i >= len(selectors) {
				break
			}
			if selectors[i] {
				result <- s
			}
		}
	}()
	return result
}

func Count(start, step int) chan int {
	// Caller should close the returned chan
	result := make(chan int)
	go func() {
		defer func() {
			// recover when the channel is closed
			recover()
		}()
		for i := start; ; i += step {
			select {
			case result <- i:
			case <-result: // 接收到通道的关闭信号
				return
			}
		}
	}()
	return result
}

func CountToClosure(start, step int) func() int {
	return func() int {
		i := start
		defer func() {
			start += step
		}()
		return i
	}
}

func Cycle(values []int) chan int {
	// Caller should close the returned chan
	r := make(chan int)
	go func() {
		defer func() {
			// recover when the channel is closed
			recover()
		}()
		for {
			for _, v := range values {
				select {
				case r <- v:
				case <-r: // 接收到通道的关闭信号
					return
				}
			}
		}
	}()
	return r
}

func CycleToClosure[T any](values []T) func() T {
	i := 0
	return func() T {
		v := values[i]
		i = (i + 1) % len(values)
		return v
	}
}

func DropWhile[T any](f func(T) bool, values []T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		drop := true
		for _, v := range values {
			if drop && f(v) {
				continue
			}
			drop = false
			ch <- v
		}
	}()
	return ch
}

func DropWhileToSlice[T any](f func(T) bool, values []T) []T {
	for i, v := range values {
		if !f(v) {
			return values[i:]
		}
	}
	return []T{}
}

func FilterFalse[T any](f func(T) bool, values []T) chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, v := range values {
			if !f(v) {
				ch <- v
			}
		}
	}()
	return ch
}

func EmptyChan[T any](ch chan T) {
	for range ch {
	}
}

func EmptyChanFunc[T any](ch chan T, f func(T)) {
	for e := range ch {
		f(e)
	}
}

func Islice[T any](slice []T, start, stop, step int) chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for i := start; i < stop && i < len(slice); i += step {
			ch <- slice[i]
		}
	}()
	return ch
}

func Islice1[T any](slice []T, start, stop, step int) chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		// for i := start; i < stop && i < len(slice); i += step {
		// 	ch <- slice[i]
		// }
		for _, v := range slice[start:stop:step] {
			ch <- v
		}
	}()
	return ch
}

func Pairwise[T any](slice []T) chan [2]T {
	ch := make(chan [2]T)
	go func() {
		defer close(ch)
		for i := 1; i < len(slice); i++ {
			ch <- [2]T{slice[i-1], slice[i]}
		}
	}()
	return ch
}

func Permutations[T any](slice []T, r int) chan []T {
	ch := make(chan []T)
	go func() {
		defer close(ch)
		permute(slice, r, ch)
	}()
	return ch
}

func permute[T any](slice []T, r int, ch chan []T) {
	if r == 0 {
		ch <- []T{}
		return
	}

	for i := 0; i < len(slice); i++ {
		// 从切片中取出第 i 个元素
		elem := slice[i]

		// 生成剩余元素的切片
		remaining := make([]T, len(slice)-1)
		copy(remaining[:i], slice[:i])
		copy(remaining[i:], slice[i+1:])

		// 递归生成剩余元素的排列
		subCh := make(chan []T)
		go func() {
			defer close(subCh)
			permute(remaining, r-1, subCh)
		}()

		// 将当前元素和子排列组合起来
		for subPerm := range subCh {
			ch <- append([]T{elem}, subPerm...)
		}
	}
}

func Repeat[T any](o T, repeat int) chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for i := 0; i < repeat; i++ {
			ch <- o
		}
	}()
	return ch
}

func RepeatToSlice[T any](o T, repeat int) []T {
	slice := make([]T, repeat)
	for i := 0; i < repeat; i++ {
		slice[i] = o
	}
	return slice
}

func RepeatSlice[T any](o []T, repeat int) chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		size := repeat * len(o)
		length := len(o)
		for i := 0; i < size; i++ {
			ch <- o[i%length]
		}
	}()
	return ch
}

func RepeatSliceToSlice[T any](o []T, repeat int) []T {
	size := repeat * len(o)
	length := len(o)
	slice := make([]T, size)
	for i := 0; i < size; i++ {
		slice[i] = o[i%length]
	}
	return slice
}

func ChanToSlice[T any](ch chan T) []T {
	r := make([]T, 0)
	for e := range ch {
		r = append(r, e)
	}
	return r
}

func SliceToChan[T any](slice []T) chan T {
	ch := make(chan T)
	go func() {
		for _, v := range slice {
			ch <- v
		}
	}()
	return ch
}

func Product[T any](slice [][]T, repeat int) chan []T {
	out := make(chan []T)
	if repeat <= 0 {
		go func() {
			out <- []T{}
			close(out)
		}()
		return out
	}
	prefix := make([]T, repeat*len(slice))
	var product func(int, int)
	product = func(i, j int) {
		// i 是第几次重复
		// j 是slice中的第几个元素
		if j == len(slice) {
			i, j = i+1, 0
		}
		if i == repeat {
			out <- append([]T{}, prefix...)
			return
		}
		for _, v := range slice[j] {
			prefix[i*len(slice)+j] = v
			product(i, j+1)
		}
	}
	go func() {
		defer close(out)
		product(0, 0)
	}()
	return out
}

func Takewhile[T any](predicate func(T) bool, slice []T) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, v := range slice {
			if !predicate(v) {
				return
			}
			out <- v
		}
	}()
	return out
}

func Tee[T any](slice []T, n int) []chan T {
	r := make([]chan T, n)
	for i := 0; i < n; i++ {
		r[i] = make(chan T)
		go func(i int) {
			defer close(r[i])
			for i2 := range slice {
				r[i] <- slice[i2]
			}
		}(i)
	}
	return r
}
