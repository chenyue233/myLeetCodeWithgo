package sort

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2
	left := MergeSort(data[:middle])
	right := MergeSort(data[middle:])
	// 合[并]
	return merge(left, right)
}
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		// 从小到大排序.
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}