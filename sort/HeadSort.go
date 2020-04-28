package sort

func initHead(nums []int,parent,len int)  {
	tem := nums[parent]
	child := 2*parent + 1
	for child < len{
		if child +1 < len && nums[child] < nums[child+1]{
			child++
		}
		if child < len && nums[child] <= tem{
			break
		}
		nums[parent] = nums[child]
		parent = child
		child = child*2 +1
	}
	nums[parent] = tem
}
func HeadSort(nums []int){
	for i := len(nums)/2; i>=0; i-- {
		initHead(nums,i,len(nums))
	}
	
	for i := len(nums)-1;i >0; i--{
		nums[0],nums[i] = nums[i],nums[0]
		
		initHead(nums,0,i)
	}
}
