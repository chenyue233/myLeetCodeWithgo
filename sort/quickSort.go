package sort

func QuickSort(nums []int,left int, right int) []int  {
	if left >= right{
		return nums
	}
	temp := nums[left]
	start := left
	stop := right
	for right != left{
		for right > left && nums[right] >= temp  {
			right --
		}
		for left < right && nums[left] <= temp  {
			left ++
		}
		if right > left{
			nums[right], nums[left] = nums[left], nums[right]
		}
	}
	nums[right], nums[start] = temp, nums[right]
	QuickSort(nums,start,left)
	QuickSort(nums,right+1,stop)
	return nums
}
