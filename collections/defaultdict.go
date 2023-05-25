package collections

type DefaultDict[U comparable, T any] struct {
	Dict[U, T]
	DefaultFactory func() T
}

func NewDefaultDict[U comparable, T any](defaultFactory func() T) *DefaultDict[U, T] {
	d, _ := NewDict[U, T]().(*Dict[U, T])
	return &DefaultDict[U, T]{
		Dict:           *d,
		DefaultFactory: defaultFactory,
	}
}

func (dd *DefaultDict[U, T]) GetItem(key U) (r T, err error) {
	if dd.Contains(key) {
		return dd.Dict.GetItem(key)
	}
	if dd.DefaultFactory != nil {
		return dd.DefaultFactory(), nil
	}
	return r, ErrKeyError
}
