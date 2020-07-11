package sort

func InsertionSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}
	for i := 1; i < n; i++ {
		if a[i] < a[i-1] {
			j := i - 1
			temp := a[i]
			for j >= 0 && a[j] > temp {
				a[j+1] = a[j]
				j--
			}
			a[j+1] = temp
		}
	}
	return a
}

func InsertionSorts(a []int, lens int) []int {
	if lens <= 1 {
		return a
	}
	for i := 1; i < lens-1; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
	return a
}
