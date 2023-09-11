package main

import "fmt"
// xoa phan tu trung nhau trong mang
func removeDuplicates(nums []int) int {
	var k int 
	for i:=0; i<len(nums); i++{
		if nums[i]!= nums[k]{
			k++
			nums[k] = nums[i]
		}
	}
	return k+1
}
func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))

}
