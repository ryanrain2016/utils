package utils

func SliceSub[T comparable](s1, s2 []T) (r []T) {
	// return slice contains element which is in s1 but not in s2
	s2Map := make(map[T]bool)
	for _, v := range s2 {
		s2Map[v] = true
	}
	rMap := make(map[T]bool)
	r = make([]T, 0, len(s1))
	for _, v := range s1 {
		if !s2Map[v] && !rMap[v] {
			r = append(r, v)
			rMap[v] = true
		}
	}
	return
}

func SliceXor[T comparable](s1, s2 []T) (r []T) {
	// return slice contains element which is in s1 and also in s2
	s2Map := make(map[T]bool)
	for _, v := range s2 {
		s2Map[v] = true
	}
	rMap := make(map[T]bool)
	r = make([]T, 0)
	for _, v := range s1 {
		if s2Map[v] && !rMap[v] {
			r = append(r, v)
			rMap[v] = true
		}
	}
	return
}

func SliceAnd[T comparable](s1, s2 []T) (r []T) {
	// return slice contains element which is in s1 or in s2
	rMap := make(map[T]bool)
	r = make([]T, 0, len(s1)+len(s2))
	for _, v := range s1 {
		if !rMap[v] {
			r = append(r, v)
			rMap[v] = true
		}
	}
	for _, v := range s2 {
		if !rMap[v] {
			r = append(r, v)
			rMap[v] = true
		}
	}
	return
}

func SliceRemoveRepeat[T comparable](s []T) []T {
	seen := make(map[T]struct{})
	rslt := make([]T, 0, len(s))

	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			rslt = append(rslt, v)
		}
	}

	return rslt
}
