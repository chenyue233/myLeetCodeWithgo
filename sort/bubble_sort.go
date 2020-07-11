package sort

func Bubble(array []int, n int) []int {
	if n <= 1 {
		return array
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if array[j+1] < array[j] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}
	return array
}
