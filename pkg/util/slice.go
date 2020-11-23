package util

func IntSliceContainsItem(array []int, item int) bool {
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

func RemoveSliceRepeatElement(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}
	set := map[int64]struct{}{}
	result := make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		if _, ok := set[arr[i]]; !ok {
			set[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}
