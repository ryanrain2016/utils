package collections

type OrderedDict[U comparable, T any] struct {
	Dict[U, T]
	keys *List[U]
}

func NewOrderedDict[U comparable, T any](o ...any) *OrderedDict[U, T] {
	dict := &OrderedDict[U, T]{
		Dict: Dict[U, T]{
			m: make(map[U]T),
		},
		keys: NewList[U](0),
	}
	for _, v := range o {
		dict.Update(v)
	}
	return dict
}

func (od *OrderedDict[U, T]) Update(m any) error {
	switch val := m.(type) {
	case IMapping[U, T]:
		for k, v := range val.Map() {
			od.SetItem(k, v)
		}
	case map[U]T:
		for k, v := range val {
			od.SetItem(k, v)
		}
	case []*MapItem[U, T]:
		for _, item := range val {
			od.SetItem(item.Key, item.Value)
		}
	case *MapItem[U, T]:
		od.SetItem(val.Key, val.Value)
	default:
		return ErrUnsupportType
	}
	return nil
}

func (od *OrderedDict[U, T]) UpdatePairSlice(keys []U, values []T) {
	var zeroT T
	for i, k := range keys {
		if !od.Contains(k) {
			od.keys.Append(k)
		}
		if i < len(values) {
			od.m[k] = values[i]
		} else {
			od.m[k] = zeroT
		}
	}
}

func (od *OrderedDict[U, T]) Items() []*MapItem[U, T] {
	items := make([]*MapItem[U, T], 0, len(od.m))
	for _, k := range od.keys.ToSlice() {
		items = append(items, &MapItem[U, T]{Key: k, Value: od.m[k]})
	}
	return items
}

func (od *OrderedDict[U, T]) PopItem() (item *MapItem[U, T], err error) {
	key, err := od.keys.Pop(-1)
	if err != nil {
		return
	}
	v, _ := od.Dict.Pop(key)
	return &MapItem[U, T]{
		Key:   key,
		Value: v,
	}, nil
}

func (od *OrderedDict[U, T]) PopFirst() (item *MapItem[U, T], err error) {
	key, err := od.keys.Pop(0)
	if err != nil {
		return
	}
	v, _ := od.Dict.Pop(key)
	return &MapItem[U, T]{
		Key:   key,
		Value: v,
	}, nil
}

func (od *OrderedDict[U, T]) SetDefault(key U, def ...T) (r T) {
	r, ok := od.m[key]
	if !ok && len(def) > 0 {
		r = def[0]
	}
	if !ok {
		od.m[key] = r
		od.keys.Append(key)
	}
	return r
}

func (od *OrderedDict[U, T]) Keys() []U {
	return od.keys.ToSlice()
}

func (od *OrderedDict[U, T]) SetItem(key U, value T) {
	if !od.Contains(key) {
		od.keys.Append(key)
	}
	od.Dict.SetItem(key, value)
}

func (od *OrderedDict[U, T]) Delete(key U) {
	od.keys.Remove(key)
	od.Dict.Delete(key)
}

func (od *OrderedDict[U, T]) Clear() {
	od.keys.Clear()
	od.Dict.Clear()
}

func (od *OrderedDict[U, T]) FromKeys(keys []U, value ...T) {
	od.Dict.FromKeys(keys, value...)
	od.keys.Clear()
	seen := NewSet[U]()
	for _, v := range keys {
		if !seen.Contains(v) {
			seen.Add(v)
			od.keys.Append(v)
		}
	}
}
