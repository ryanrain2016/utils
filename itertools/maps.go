package itertools

func MapCopy[U comparable, T any](m map[U]T) map[U]T {
	c := make(map[U]T)
	for k := range m {
		c[k] = m[k]
	}
	return c
}
