package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6},9))
}
func searchInsert(nums []int, target int) int {
	if nums==nil || len(nums)==0{
		return -1
	}
	left ,right :=0,len(nums)
	for left <right{
		mid :=left+(right-left)/2
		if nums[mid] >= target{
			right =mid
		}else if  nums[mid]<target{
			left =mid+1
		}
	}
	return left
}