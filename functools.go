package utils

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
