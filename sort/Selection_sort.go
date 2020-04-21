package sort

func Selection(n []int) (res []int) {
	a := len(n)
	if a <= 1 {
		return n
	}
	for i:=0;i<a;i++{
		minIndex := i
		for j := i+1;j<a;j++{
			if n[j] <n[minIndex]{
				minIndex = j
			}
		}
		n[i],n[minIndex] = n[minIndex],n[i]
	}
	return n
}
