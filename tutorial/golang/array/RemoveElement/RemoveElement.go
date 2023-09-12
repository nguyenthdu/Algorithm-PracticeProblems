package main

import "fmt"

func removeElement(nums []int, val int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]//counts = 0, i = 1, val i =2, ,phan tu i se duoc ghi de len phan tu count
			count++
		}
	}
	return count
}
func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}
