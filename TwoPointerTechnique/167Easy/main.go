package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{2,7,11,15},18))
}


func twoSum(numbers []int, target int) []int {
	if len(numbers)<2 || numbers==nil{
		return nil
	}
	left,right :=0,len(numbers)-1
	for left<right{
		sum :=numbers[left]+numbers[right]
		if sum==target{
			return []int{left+1,right+1}
		}else if sum>target{
			right--
		}else {
			left++
		}
	}
	return nil
}