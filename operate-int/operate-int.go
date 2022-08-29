package tool1

type Int struct{}

func NewInt() Int {
	return Int{}
}

// int 类型数据的操作
func (t Int) Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (t Int) Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (t Int) Swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}
