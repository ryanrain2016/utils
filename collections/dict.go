package collections

import "reflect"

type MapItem[U comparable, T any] struct {
	Key   U
	Value T
}

type IMapping[U comparable, T any] interface {
	Map() map[U]T                  // 返回内部的mao
	Copy() IMapping[U, T]          // 返回一个拷贝，拷贝主要是内部map
	FromKeys(keys []U, value ...T) // 通过K创建dict，value没有传取0值，否则取第一个value值，后面的参数丢掉
	Get(key U, defs ...T) (r T)    // 获取key对应的值，defs为默认值，没有key时返回，没有传defs则返回0值
	GetItem(key U) (r T, err error)
	SetItem(key U, value T)
	Items() []*MapItem[U, T]                   // 遍历元素返回item组成的slice
	Keys() []U                                 // 返回key组成的slice
	Pop(key U, defs ...T) (t T, err error)     // should delete popped key
	PopItem() (item *MapItem[U, T], err error) // should delete popped item
	SetDefault(key U, def ...T) (r T)          // 类似python中dict.setdefault的功能，key存在时返回对应的值，否则设置key对应的值为def，并返回，def不传时返回0值
	UpdatePairSlice(keys []U, values []T)
	Update(m any) error
	Len() int
	Delete(key U)
	Contains(key U) bool
	Clear()
	Eq(other IMapping[U, T]) bool
}

type Dict[U comparable, T any] struct {
	m map[U]T
}

func NewDict[U comparable, T any](o ...any) IMapping[U, T] {
	dict := &Dict[U, T]{
		m: make(map[U]T),
	}
	for _, v := range o {
		dict.Update(v)
	}
	return dict
}

func (d *Dict[U, T]) Len() int {
	return len(d.m)
}

func (d *Dict[U, T]) Clear() {
	d.m = make(map[U]T)
}

func (d *Dict[U, T]) Map() map[U]T {
	return d.m
}

func (d *Dict[U, T]) Copy() IMapping[U, T] {
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
	d.m = m
}

func (d *Dict[U, T]) Get(key U, defs ...T) (r T) {
	r, ok := d.m[key]
	if !ok && len(defs) > 0 {
		r = defs[0]
	}
	return r
}

func (d *Dict[U, T]) GetItem(key U) (r T, err error) {
	r, ok := d.m[key]
	if ok {
		return r, nil
	}
	return r, ErrKeyError
}

func (d *Dict[U, T]) SetItem(key U, value T) {
	d.m[key] = value
}

func (d *Dict[U, T]) Contains(key U) bool {
	_, ok := d.m[key]
	return ok
}

func (d *Dict[U, T]) Delete(key U) {
	delete(d.m, key)
}

func (d *Dict[U, T]) Items() []*MapItem[U, T] {
	items := make([]*MapItem[U, T], 0, len(d.m))
	for k, v := range d.m {
		items = append(items, &MapItem[U, T]{Key: k, Value: v})
	}
	return items
}

func (d *Dict[U, T]) Keys() []U {
	keys := make([]U, 0, len(d.m))
	for k := range d.m {
		keys = append(keys, k)
	}
	return keys
}

func (d *Dict[U, T]) Pop(key U, defs ...T) (t T, err error) {
	t, ok := d.m[key]
	if !ok && len(defs) > 0 {
		t = defs[0]
	}
	delete(d.m, key)
	if !ok {
		err = ErrKeyError
	}
	return t, err
}

func (d *Dict[U, T]) PopItem() (item *MapItem[U, T], err error) {
	if len(d.m) == 0 {
		err = ErrPopFromEmptyCollections
		return nil, err
	}
	for k, v := range d.m {
		delete(d.m, k)
		item = &MapItem[U, T]{Key: k, Value: v}
		break
	}
	return item, nil
}

func (d *Dict[U, T]) SetDefault(key U, def ...T) (r T) {
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
	switch val := m.(type) {
	case IMapping[U, T]:
		for k, v := range val.Map() {
			d.m[k] = v
		}
	case map[U]T:
		for k, v := range val {
			d.m[k] = v
		}
	case []*MapItem[U, T]:
		for _, item := range val {
			d.m[item.Key] = item.Value
		}
	case *MapItem[U, T]:
		d.m[val.Key] = val.Value
	default:
		return ErrUnsupportType
	}
	return nil
}

func (d *Dict[U, T]) Eq(other IMapping[U, T]) bool {
	return reflect.DeepEqual(d.Map(), other.Map())
}
