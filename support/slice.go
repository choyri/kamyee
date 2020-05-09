package support

func InStringSlice(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func InInt64Slice(needle int64, haystack []int64) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func InUint64Slice(needle uint64, haystack []uint64) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func InIntSlice(needle int, haystack []int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func InUintSlice(needle uint, haystack []uint) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func IntersectStringSlice(a []string, b []string) []string {
	var (
		tmpMap = make(map[string]struct{})
		ret    = make([]string, 0)
	)

	for _, v := range a {
		tmpMap[v] = struct{}{}
	}

	for _, v := range b {
		if _, exists := tmpMap[v]; exists {
			ret = append(ret, v)
		}
	}

	return ret
}

func IntersectInt64Slice(a []int64, b []int64) []int64 {
	var (
		tmpMap = make(map[int64]struct{})
		ret    = make([]int64, 0)
	)

	for _, v := range a {
		tmpMap[v] = struct{}{}
	}

	for _, v := range b {
		if _, exists := tmpMap[v]; exists {
			ret = append(ret, v)
		}
	}

	return ret
}

func IntersectUint64Slice(a []uint64, b []uint64) []uint64 {
	var (
		tmpMap = make(map[uint64]struct{})
		ret    = make([]uint64, 0)
	)

	for _, v := range a {
		tmpMap[v] = struct{}{}
	}

	for _, v := range b {
		if _, exists := tmpMap[v]; exists {
			ret = append(ret, v)
		}
	}

	return ret
}

func IntersectIntSlice(a []int, b []int) []int {
	var (
		tmpMap = make(map[int]struct{})
		ret    = make([]int, 0)
	)

	for _, v := range a {
		tmpMap[v] = struct{}{}
	}

	for _, v := range b {
		if _, exists := tmpMap[v]; exists {
			ret = append(ret, v)
		}
	}

	return ret
}

func IntersectUintSlice(a []uint, b []uint) []uint {
	var (
		tmpMap = make(map[uint]struct{})
		ret    = make([]uint, 0)
	)

	for _, v := range a {
		tmpMap[v] = struct{}{}
	}

	for _, v := range b {
		if _, exists := tmpMap[v]; exists {
			ret = append(ret, v)
		}
	}

	return ret
}

func GetAddedAndDeletedInUintSlice(old, new []uint) (added, deleted []uint) {
	var (
		oldMap = make(map[uint]struct{})
		newMap = make(map[uint]struct{})
	)

	for _, v := range old {
		oldMap[v] = struct{}{}
	}

	for _, v := range new {
		newMap[v] = struct{}{}
		if _, exists := oldMap[v]; !exists {
			added = append(added, v)
		}
	}

	for _, v := range old {
		if _, exists := newMap[v]; !exists {
			deleted = append(deleted, v)
		}
	}

	return
}
