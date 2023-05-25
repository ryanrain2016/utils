package safe

import (
	"reflect"
	"sync"

	"github.com/ryanrain2016/utils/collections"
)

type Dict[U comparable, T any] struct {
	m map[U]T
	sync.RWMutex
}

func NewDict[U comparable, T any](items ...map[U]T) collections.IMapping[U, T] {
	dict := &Dict[U, T]{
		m: make(map[U]T),
	}
	for _, v := range items {
		for k, v2 := range v {
			dict.m[k] = v2
		}
	}
	return dict
}

func (d *Dict[U, T]) Len() int {
	d.RLock()
	defer d.RUnlock()
	return len(d.m)
}

func (d *Dict[U, T]) Clear() {
	d.Lock()
	defer d.Unlock()
	d.m = make(map[U]T)
}

func (d *Dict[U, T]) Map() map[U]T {
	d.RLock()
	defer d.RUnlock()
	return d.m
}

func (d *Dict[U, T]) Copy() collections.IMapping[U, T] {
	d.RLock()
	defer d.RUnlock()
	cpy := make(map[U]T, len(d.m))
	for k, v := range d.m {
		cpy[k] = v
	}
	return &Dict[U, T]{m: cpy}
}

func (d *Dict[U, T]) FromKeys(keys []U, value ...T) {
	m := make(map[U]T)
	for _, k := range keys {
		var v T
		if len(value) > 0 {
			v = value[0]
		}
		m[k] = v
	}
	d.Lock()
	defer d.Unlock()
	d.m = m
}

func (d *Dict[U, T]) Get(key U, defs ...T) (r T) {
	d.RLock()
	defer d.RUnlock()
	r, ok := d.m[key]
	if !ok && len(defs) > 0 {
		r = defs[0]
	}
	return r
}

func (d *Dict[U, T]) GetItem(key U) (r T, err error) {
	d.RLock()
	defer d.RUnlock()
	r, ok := d.m[key]
	if ok {
		return r, nil
	}
	return r, collections.ErrIndexOutofRange
}

func (d *Dict[U, T]) SetItem(key U, value T) {
	d.Lock()
	defer d.Unlock()
	d.m[key] = value
}

func (d *Dict[U, T]) Contains(key U) bool {
	d.Lock()
	defer d.Unlock()
	_, ok := d.m[key]
	return ok
}

func (d *Dict[U, T]) Delete(key U) {
	d.Lock()
	defer d.Unlock()
	delete(d.m, key)
}

func (d *Dict[U, T]) Items() []*collections.MapItem[U, T] {
	d.RLock()
	defer d.RUnlock()
	items := make([]*collections.MapItem[U, T], 0, len(d.m))
	for k, v := range d.m {
		items = append(items, &collections.MapItem[U, T]{Key: k, Value: v})
	}
	return items
}

func (d *Dict[U, T]) Keys() []U {
	d.RLock()
	defer d.RUnlock()
	keys := make([]U, 0, len(d.m))
	for k := range d.m {
		keys = append(keys, k)
	}
	return keys
}

func (d *Dict[U, T]) Pop(key U, defs ...T) (t T, err error) {
	d.Lock()
	defer d.Unlock()
	t, ok := d.m[key]
	if !ok && len(defs) > 0 {
		t = defs[0]
	}
	delete(d.m, key)
	if !ok {
		err = collections.ErrKeyError
	}
	return t, err
}

func (d *Dict[U, T]) PopItem() (item *collections.MapItem[U, T], err error) {
	d.Lock()
	defer d.Unlock()
	if len(d.m) == 0 {
		err = collections.ErrPopFromEmptyCollections
		return nil, err
	}
	for k, v := range d.m {
		delete(d.m, k)
		item = &collections.MapItem[U, T]{Key: k, Value: v}
		break
	}
	return item, nil
}

func (d *Dict[U, T]) SetDefault(key U, def ...T) (r T) {
	d.Lock()
	defer d.Unlock()
	r, ok := d.m[key]
	if !ok && len(def) > 0 {
		r = def[0]
	}
	if !ok {
		d.m[key] = r
	}
	return r
}

func (d *Dict[U, T]) UpdatePairSlice(keys []U, values []T) {
	d.Lock()
	defer d.Unlock()
	var zeroT T
	for i, k := range keys {
		if i < len(values) {
			d.m[k] = values[i]
		} else {
			d.m[k] = zeroT
		}
	}
}

func (d *Dict[U, T]) Update(m any) error {
	d.Lock()
	defer d.Unlock()

	switch val := m.(type) {
	case collections.IMapping[U, T]:
		for k, v := range val.Map() {
			d.m[k] = v
		}
	case map[U]T:
		for k, v := range val {
			d.m[k] = v
		}
	case []*collections.MapItem[U, T]:
		for _, item := range val {
			d.m[item.Key] = item.Value
		}
	default:
		return collections.ErrUnsupportType
	}
	return nil
}

func (d *Dict[U, T]) Eq(other collections.IMapping[U, T]) bool {
	d.RLock()
	defer d.RUnlock()
	return reflect.DeepEqual(d.Map(), other.Map())
}
