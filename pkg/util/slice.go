package util

func IntSliceContainsItem(array []int, item interface{}) bool {
	if array == nil || len(array) == 0 {
		return false
	}
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return true
		}
	}
	return false
}
