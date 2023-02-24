package utils

func InSlice(a []string, x string) bool {
	for _, v := range a {
		if x == v {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, v := range a {
		if x == v {
			return true
		}
	}
	return false
}
