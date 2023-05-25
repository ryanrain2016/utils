package functools

import (
	"errors"
	"math"
	"reflect"

	"github.com/ryanrain2016/utils"
	"github.com/ryanrain2016/utils/itertools"
	"github.com/ryanrain2016/utils/types"
)

// 将两个值相加
func Add[T types.Addable](a, b T) T {
	return a + b
}

// 将两个值相减
func Sub[T types.Number](a, b T) T {
	return a - b
}

// 将两个值相乘
func Mul[T types.Number](a, b T) T {
	return a * b
}

// 将两个值相除
func TrueDiv[T types.OrderedNumber](a, b T) float64 {
	return float64(a) / float64(b)
}

// 将两个值相除并向下取整
func FloorDiv[T types.Int](a, b T) T {
	return a / b
}

// 计算两个值的模数
func Mod[T types.Int](a, b T) T {
	return a % b
}

// 将一个值的幂次方
func Pow[T types.OrderedNumber](a, b T) T {
	return T(math.Pow(float64(a), float64(b)))
}

// 比较两个值是否小于
func Lt[T types.Ordered](a, b T) bool {
	return a < b
}

// 比较两个值是否小于或等于
func Le[T types.Ordered](a, b T) bool {
	return a <= b
}

// 比较两个值是否相等
func Eq[T comparable](a, b T) bool {
	return a == b
}

// 比较两个值是否不相等
func Ne[T comparable](a, b T) bool {
	return a != b
}

// 比较两个值是否大于或等于
func Ge[T types.Ordered](a, b T) bool {
	return a >= b
}

// 比较两个值是否大于
func Gt[T types.Ordered](a, b T) bool {
	return a > b
}

type booler interface {
	Bool() bool
}

type stringer interface {
	String() string
}

type lener interface {
	Len() int
}

type sizer interface {
	Size() int
}

// Truth 返回常见类型的真值
func Truth(o any) (t bool) {
	defer func() {
		e := recover()
		if e != nil {
			t = false
		}
	}()
	switch v := o.(type) {
	case int:
		return v != 0
	case float32:
		return v != 0
	case float64:
		return v != 0
	case int8:
		return v != 0
	case int16:
		return v != 0
	case int32:
		return v != 0
	case int64:
		return v != 0
	case uint:
		return v != 0
	case uint8:
		return v != 0
	case uint16:
		return v != 0
	case uint32:
		return v != 0
	case uint64:
		return v != 0
	case complex64:
		return v != 0
	case complex128:
		return v != 0
	case string:
		return v != ""
	default:
		if v == nil {
			return false
		}
		if obj, ok := v.(booler); ok {
			return obj.Bool()
		}
		if obj, ok := v.(lener); ok {
			return obj.Len() != 0
		}
		if obj, ok := v.(sizer); ok {
			return obj.Size() != 0
		}
		if obj, ok := v.(stringer); ok {
			if obj.String() == "" {
				return false
			}
		}
		objType := reflect.TypeOf(v)
		objValue := reflect.ValueOf(v)
		if objType.Kind() == reflect.Ptr {
			if objValue.IsNil() {
				return false
			} else {
				objValue = objValue.Elem()
			}
		}
		if !objValue.IsValid() {
			return false
		}
		if objValue.Kind() == reflect.Struct {
			return true
		}
		if objValue.IsZero() {
			return false
		}
		if objType.Kind() == reflect.Array ||
			objType.Kind() == reflect.Slice ||
			objType.Kind() == reflect.Map ||
			objType.Kind() == reflect.Chan {
			return objValue.Len() != 0
		}
		return true
	}
}

func Not(o any) bool {
	return !Truth(o)
}

func Is(a any, b any) bool {
	return &a == &b
}

func IsNot(a any, b any) bool {
	return !Is(a, b)
}

func And[T types.Int](a, b T) T {
	return a & b
}

type indexer interface {
	Index() int
}

func Index(a any) int {
	if v, ok := a.(indexer); ok {
		return v.Index()
	}
	return -1
}

func Inv[T types.Int](a T) T {
	return ^a
}

func Invert[T types.Int](a T) T {
	return ^a
}

func Lshift[T types.Int](a, b T) T {
	return a << b
}

func Neg[T types.Number](a T) T {
	return -a
}

func Or[T types.Int](a, b T) T {
	return a | b
}

func Rshift[T types.Int](a, b T) T {
	return a >> b
}

func Xor[T types.Int](a, b T) T {
	return a ^ b
}

func Concat[T any](a, b []T) []T {
	return append(append([]T(nil), a...), b...)
}

func Contains[T comparable](a []T, b T) bool {
	for _, v := range a {
		if b == v {
			return true
		}
	}
	return false
}

func Container[T comparable](a []T) func(T) bool {
	return func(t T) bool {
		return Contains(a, t)
	}
}

func MapValueContains[T, U comparable](a map[U]T, b T) bool {
	for _, v := range a {
		if b == v {
			return true
		}
	}
	return false
}

func MapValueContainer[T, U comparable](a map[U]T) func(T) bool {
	return func(t T) bool {
		return MapValueContains(a, t)
	}
}

func MapKeyContains[T, U comparable](a map[U]T, b U) bool {
	_, ok := a[b]
	return ok
}

func MapKeyContainer[T, U comparable](a map[U]T) func(b U) bool {
	return func(t U) bool {
		return MapKeyContains(a, t)
	}
}

func CountOf[T comparable](a []T, b T) (r int) {
	for _, v := range a {
		if b == v {
			r++
		}
	}
	return
}

func IndexOf[T comparable](a []T, b T) int {
	return itertools.FindIndex(a, b)
}

func GetItem[T any, U comparable](a map[U]T, key U) T {
	value := a[key]
	return value
}

func DeleteItem[T any, U comparable](a map[U]T, key U) {
	delete(a, key)
}

func SetItem[T any, U comparable](a map[U]T, key U, value T) {
	a[key] = value
}

// 判断一个值是否是结构体的指针
func isPointer(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Ptr
}

// 判断一个值是否是指向结构体的指针
func isPointerToStruct(v any) bool {
	// 如果不是指针类型，返回 false
	if !isPointer(v) {
		return false
	}
	// 获取指针指向的对象的类型
	t := reflect.TypeOf(v).Elem()
	// 判断指向的对象是否是结构体类型
	return t.Kind() == reflect.Struct
}

// 判断一个值是否是结构体
func isStruct(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Struct
}

func GetField[T any](a any, field string) (r T) {
	var v = reflect.ValueOf(a)
	if isPointerToStruct(a) {
		v = v.Elem()
	} else if !isStruct(a) {
		return
	}
	fieldValue := v.FieldByName(field)
	if fieldValue.IsValid() {
		o := fieldValue.Interface()
		return o.(T)
	}
	return
}

func SetField(a any, field string, value any) error {
	if !isPointerToStruct(a) {
		return errors.New("only can set field of a pointer to a struct")
	}
	fieldValue := reflect.ValueOf(a).Elem().FieldByName(field)
	if !fieldValue.IsValid() {
		return errors.New("field not exist")
	}
	fieldValue.Set(reflect.ValueOf(value))
	return nil
}

func FieldGetter[T any](field string) func(a any) T {
	return func(a any) T {
		return GetField[T](a, field)
	}
}

func FieldSetter(field string, value any) func(a any) {
	return func(a any) {
		SetField(a, field, value)
	}
}

func ItemGetter[T any, U comparable](key U) func(map[U]T) T {
	return func(m map[U]T) T {
		return GetItem(m, key)
	}
}

func CallMethod(a any, name string, args ...any) (r []any, err error) {
	var v = reflect.ValueOf(a)
	if isPointerToStruct(a) {
		v = v.Elem()
	} else if !isStruct(a) {
		err = errors.New("CallMethod should be called on struct")
		return
	}
	method := v.MethodByName(name)
	if !method.IsValid() {
		err = errors.New("method not exist")
		return
	}
	rslt := method.Call(utils.Map(func(arg any) reflect.Value { return reflect.ValueOf(arg) }, args))
	return utils.Map(func(v reflect.Value) any { return v.Interface() }, rslt), nil
}

func MethodCaller(name string, args ...any) func(a any) (r []any, err error) {
	return func(a any) (r []any, err error) {
		return CallMethod(a, name, args...)
	}
}
