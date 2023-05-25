package collections

import (
	"reflect"
	"sort"

	"github.com/ryanrain2016/utils/itertools"
)

type Counter[T comparable] struct {
	m     map[T]int
	items []T
}

func NewCounter[T comparable]() *Counter[T] {
	return &Counter[T]{
		m:     make(map[T]int),
		items: make([]T, 0),
	}
}

func (c *Counter[T]) Contains(el T) bool {
	_, ok := c.m[el]
	return ok
}

func (c *Counter[T]) Clear() {
	c.items = make([]T, 0)
	c.m = make(map[T]int)
}

func (c *Counter[T]) Delete(el T) {
	index := itertools.FindIndex(c.items, el)
	c.items = append(c.items[:index], c.items[index+1:]...)
	delete(c.m, el)
}

func (c *Counter[T]) FromSlice(items []T) {
	m := make(map[T]int)
	items_ := make([]T, 0)
	for _, v := range items {
		if _, ok := m[v]; !ok {
			items_ = append(items_, v)
		}
		m[v]++
	}
	*c = Counter[T]{
		items: items_,
		m:     m,
	}
}

func (c *Counter[T]) FromMap(m map[T]int) {
	items := make([]T, 0)
	for k := range m {
		items = append(items, k)
	}
	*c = Counter[T]{
		items: items,
		m:     itertools.MapCopy(m),
	}
}

func (c *Counter[T]) Copy() IMapping[T, int] {
	ct := NewCounter[T]()
	ct.m = itertools.MapCopy(c.m)
	ct.items = append(ct.items, c.items...)
	return ct
}

func (c *Counter[T]) Count(t T) {
	if _, ok := c.m[t]; !ok {
		c.items = append(c.items, t)
	}
	c.m[t]++
}

func (c *Counter[T]) Elements() []T {
	r := make([]T, 0)
	for _, el := range c.items {
		cnt := c.m[el]
		for i := 0; i < cnt; i++ {
			r = append(r, el)
		}
	}
	return r
}

func (c *Counter[T]) MostCommon(n ...uint) []*MapItem[T, int] {
	num := len(c.m)
	if len(n) > 0 {
		num = int(n[0])
	}
	items := make([]*MapItem[T, int], 0, len(c.m))
	for k, v := range c.m {
		items = append(items, &MapItem[T, int]{Key: k, Value: v})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})
	return items[:num]
}

func (c *Counter[T]) toMap(other any) map[T]int {
	var otherMap map[T]int
	if v, ok := other.(*Counter[T]); ok {
		otherMap = v.m
	}
	switch v := other.(type) {
	case *Counter[T]:
		otherMap = v.m
	case map[T]int:
		otherMap = v
	case []T:
		otherMap = make(map[T]int)
		for _, v := range v {
			otherMap[v]++
		}
	case *MapItem[T, int]:
		otherMap = make(map[T]int)
		otherMap[v.Key] = v.Value
	case []*MapItem[T, int]:
		otherMap = make(map[T]int)
		for _, mi := range v {
			otherMap[mi.Key] = mi.Value
		}
	case IMapping[T, int]:
		return v.Map()
	}
	return otherMap
}

func (c *Counter[T]) Subtract(other any) {
	for k, v := range c.toMap(other) {
		c.m[k] -= v
	}
}

func (c *Counter[T]) Total() int {
	sum := 0
	for _, v := range c.m {
		sum += v
	}
	return sum
}

func (c *Counter[T]) FromKeys([]T, ...int) {
	panic("FromKeys is not implemented for Counter objects")
}

func (c *Counter[T]) Update(other any) error {
	otherMap := c.toMap(other)
	if otherMap == nil {
		return ErrUnsupportType
	}
	for k, v := range otherMap {
		c.m[k] += v
	}
	return nil
}

// Counts returns the count of a given element in the counter.
func (c *Counter[T]) Get(elem T, defs ...int) int {
	return c.m[elem]
}

func (c *Counter[T]) GetItem(elem T) (int, error) {
	return c.m[elem], nil
}

func (c *Counter[T]) Items() []*MapItem[T, int] {
	r := make([]*MapItem[T, int], 0)
	for _, v := range c.items {
		r = append(r, &MapItem[T, int]{
			Key:   v,
			Value: c.m[v],
		})
	}
	return r
}

func (c *Counter[T]) Keys() []T {
	return c.items[:]
}

func (c *Counter[T]) Len() int {
	return len(c.items)
}

func (c *Counter[T]) Map() map[T]int {
	return itertools.MapCopy(c.m)
}

func (c *Counter[T]) Pop(key T, defs ...int) (t int, err error) {
	if c.Contains(key) {
		defer func() {
			c.Delete(key)
		}()
		return c.Get(key), nil
	}
	if len(defs) > 0 {
		return defs[0], nil
	}
	return t, ErrElementNotFound
}

func (c *Counter[T]) PopItem() (r *MapItem[T, int], err error) {
	for _, v := range c.items {
		r = &MapItem[T, int]{
			Key:   v,
			Value: c.Get(v),
		}
		c.Delete(v)
		return r, nil
	}
	return r, ErrPopFromEmptyCollections
}

func (c *Counter[T]) SetItem(key T, value int) {
	if !c.Contains(key) {
		c.items = append(c.items, key)
	}
	c.m[key] = value
}

func (c *Counter[T]) SetDefault(key T, def ...int) (r int) {
	r, ok := c.m[key]
	if !ok && len(def) > 0 {
		r = def[0]
	}
	if !ok {
		c.m[key] = r
		c.items = append(c.items, key)
	}
	return r
}

func (c *Counter[T]) UpdatePairSlice(keys []T, values []int) {
	for i, k := range keys {
		if !c.Contains(k) {
			c.items = append(c.items, k)
		}
		if i < len(values) {
			c.m[k] += values[i]
		} else {
			c.m[k] = 0
		}
	}
}

func (c *Counter[T]) Eq(other IMapping[T, int]) bool {
	return reflect.DeepEqual(c.Map(), other.Map())
}
