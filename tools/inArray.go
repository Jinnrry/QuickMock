package tools

// InArray 判断某个值是否在数组中
func InArray[T comparable](point T, arg []T) bool {
	for _, item := range arg {
		if item == point {
			return true
		}
	}
	return false
}
