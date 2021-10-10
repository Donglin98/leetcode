package main
//旋转数组的平方
func main() {
}
func rotate(nums []int, k int)  {
	le,ans :=len(nums),make([]int,len(nums))
	for i ,_ :=range nums{
		ans[i] =nums[i]
	}
	for  i,_ :=range ans{
		nums[(i+k)%le] =ans[i]
	}
}
