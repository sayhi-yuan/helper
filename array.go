package helper

// 数组去重
func ArrayUnique[T StringInteger](arr []T) []T {
	m := map[T]struct{}{}
	nArr := []T{}

	for _, item := range arr {
		if _, ok := m[item]; !ok {
			m[item] = struct{}{}
			nArr = append(nArr, item)
		}
	}

	return nArr
}

// 判断元素是否在数组中
func ArrayIn[T StringInteger](arr []T, e T) bool {
	m := map[T]struct{}{}
	for _, item := range arr {
		m[item] = struct{}{}
	}

	if _, ok := m[e]; ok {
		return true
	}

	return false
}
