package functools

import (
	"reflect"
	"strings"
)

func Reduce[T any](reduceFunc func(T, T) T, slice []T) T {
	var rslt T
	if len(slice) == 0 {
		return rslt
	}
	rslt = slice[0]
	for _, v := range slice[1:] {
		rslt = reduceFunc(rslt, v)
	}
	return rslt
}

func ReduceWithOrigin[T any, U any](reduceFunc func(T, U) T, slice []U, origin T) T {
	rslt := origin
	for _, v := range slice {
		rslt = reduceFunc(rslt, v)
	}
	return rslt
}

func FuncSignature(f any) string {
	fn := reflect.ValueOf(f)
	if fn.Kind() != reflect.Func {
		panic("key is not callable")
	}
	ft := fn.Type()
	var sb strings.Builder
	sb.WriteString("func")
	if name := ft.Name(); name != "" {
		sb.WriteString(" ")
		sb.WriteString(ft.Name())
	}
	sb.WriteString("(")
	ins := make([]string, 0)
	for i := 0; i < ft.NumIn(); i++ {
		ins = append(ins, ft.In(i).String())
	}
	sb.WriteString(strings.Join(ins, ", "))
	sb.WriteString(")")
	outs := make([]string, 0)
	for i := 0; i < ft.NumOut(); i++ {
		outs = append(outs, ft.Out(i).String())
	}
	sb.WriteString(strings.Join(outs, ", "))
	return sb.String()
}
