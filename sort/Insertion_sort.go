package sort

func InsertionSort(a []int)[]int{
    n := len(a)
    if n <= 1{
        return a
    }
    for i := 1;i <n;i++ {
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
