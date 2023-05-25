package safe

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/ryanrain2016/utils"
	"github.com/ryanrain2016/utils/collections"
)

type ChainMap[U comparable, T any] struct {
	items []collections.IMapping[U, T]
	sync.RWMutex
}

func NewChainMap[U comparable, T any](items ...collections.IMapping[U, T]) *ChainMap[U, T] {
	cm := &ChainMap[U, T]{
		items: make([]collections.IMapping[U, T], len(items)),
	}
	copy(cm.items, items)
	if len(cm.items) == 0 {
		cm.items = append(cm.items, NewDict[U, T]())
	}
	return cm
}

func (c *ChainMap[U, T]) Contains(key U) bool {
	c.RLock()
	defer c.RUnlock()
	return utils.AnyFunc(func(u collections.IMapping[U, T]) bool { return u.Contains(key) }, c.items...)
}

func (c *ChainMap[U, T]) Maps() []collections.IMapping[U, T] {
	c.RLock()
	defer c.RUnlock()
	return c.items
}

func (c *ChainMap[U, T]) NewChild(m collections.IMapping[U, T]) *ChainMap[U, T] {
	c.RLock()
	defer c.RUnlock()
	if m == nil {
		m = NewDict[U, T]()
	}
	items := []collections.IMapping[U, T]{m}
	items = append(items, c.items...)
	return NewChainMap(items...)
}

func (c *ChainMap[U, T]) Parents() *ChainMap[U, T] {
	c.RLock()
	defer c.RUnlock()
	return NewChainMap(c.items[1:]...)
}

func (c *ChainMap[U, T]) Clear() {
	c.Lock()
	defer c.Unlock()
	c.items = make([]collections.IMapping[U, T], 0)
}

func (c *ChainMap[U, T]) Copy() collections.IMapping[U, T] {
	c.RLock()
	defer c.RUnlock()
	if len(c.items) == 0 {
		return NewChainMap[U, T]()
	}
	items := make([]collections.IMapping[U, T], len(c.items))
	for i, v := range c.items {
		items[i] = v.Copy()
	}
	return NewChainMap(items...)
}

func (c *ChainMap[U, T]) FromKeys(keys []U, value ...T) {
	cm := NewChainMap[U, T]()
	var t T
	if len(value) > 0 {
		t = value[0]
	}
	for _, key := range keys {
		cm.items[0].SetItem(key, t)
	}
	c.Lock()
	defer c.Unlock()
	c.items = cm.items
}

func (c *ChainMap[U, T]) Get(key U, defs ...T) (r T) {
	c.RLock()
	defer c.RUnlock()
	for _, item := range c.items {
		if item.Contains(key) {
			return item.Get(key)
		}
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return
}

func (c *ChainMap[U, T]) GetItem(key U) (r T, err error) {
	c.RLock()
	defer c.RUnlock()
	for _, item := range c.items {
		if item.Contains(key) {
			return item.GetItem(key)
		}
	}
	return r, collections.ErrKeyError
}

func (c *ChainMap[U, T]) SetItem(key U, value T) {
	c.Lock()
	defer c.Unlock()
	c.items[0].SetItem(key, value)
}

func (c *ChainMap[U, T]) Delete(key U) {
	c.Lock()
	defer c.Unlock()
	m := c.items[0]
	m.Delete(key)
}

func (c *ChainMap[U, T]) Len() int {
	return len(c.Map())
}

func (c *ChainMap[U, T]) Map() map[U]T {
	c.RLock()
	defer c.RUnlock()
	r := make(map[U]T)
	for _, m := range c.items {
		for k, v := range m.Map() {
			if _, ok := r[k]; !ok {
				r[k] = v
			}
		}
	}
	return r
}

func (c *ChainMap[U, T]) Items() []*collections.MapItem[U, T] {
	c.RLock()
	defer c.RUnlock()
	items := make([]*collections.MapItem[U, T], 0)
	seen := NewSet[U]()
	for _, v := range c.items {
		for _, item := range v.Items() {
			if !seen.Contains(item.Key) {
				seen.Add(item.Key)
				items = append(items, item)
			}
		}
	}
	return items
}

func (c *ChainMap[U, T]) Keys() []U {
	c.RLock()
	defer c.RUnlock()
	items := make([]U, 0)
	seen := NewSet[U]()
	for _, v := range c.items {
		for _, item := range v.Items() {
			if !seen.Contains(item.Key) {
				seen.Add(item.Key)
				items = append(items, item.Key)
			}
		}
	}
	return items
}

func (c *ChainMap[U, T]) Pop(key U, defs ...T) (t T, err error) {
	c.Lock()
	defer c.Unlock()
	m := c.items[0]
	if v, err := m.Pop(key); err == nil {
		return v, nil
	}

	if len(defs) > 0 {
		return defs[0], nil
	}
	err = fmt.Errorf("key not found in the first mapping: %v", key)
	return
}

func (c *ChainMap[U, T]) PopItem() (item *collections.MapItem[U, T], err error) {
	c.Lock()
	defer c.Unlock()
	m := c.items[0]
	item, err = m.PopItem()
	if err == nil {
		return item, err
	}
	err = errors.New("no keys found in the first mapping")
	return
}

func (c *ChainMap[U, T]) SetDefault(key U, def ...T) (r T) {
	c.Lock()
	defer c.Unlock()
	m := c.items[0]
	return m.SetDefault(key, def...)
}

func (c *ChainMap[U, T]) UpdatePairSlice(keys []U, values []T) {
	c.Lock()
	defer c.Unlock()
	m := c.items[0]
	m.UpdatePairSlice(keys, values)
}

func (c *ChainMap[U, T]) Update(m any) error {
	c.Lock()
	defer c.Unlock()
	m1 := c.items[0]
	return m1.Update(m)
}

func (c *ChainMap[U, T]) Eq(other collections.IMapping[U, T]) bool {
	return reflect.DeepEqual(c.Map(), other.Map())
}
