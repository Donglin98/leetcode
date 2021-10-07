package main

import "fmt"

//704 easy
/*
	核心思想 升序定义两个首尾指针索引计算出mid值比较target
*/
func search(nums []int, target int) int {
	left ,right :=0,len(nums)-1
	for left<=right {
		mid :=left+(right-left)>>1
		if nums[mid]==target{
			return mid
		}else if nums[mid]>target {
			right=mid-1
		}
		left=mid+1
	}
	return -1
}

func main() {
	//zz :=[]int{-1,0,3,5,9,12}
	//search(zz,9)
	fmt.Println(search([]int{-1,0,3,5,9,12},9))
}
