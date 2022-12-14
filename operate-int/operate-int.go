package tool1

// int 类型数据的操作
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}
