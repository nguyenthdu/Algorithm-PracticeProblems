package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	sum := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			sum++
		}
		if nums[i] == 0 {
			sum = 0
		}
		if sum > max {
			max = sum
		}
	}
	return max

}
func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Print(findMaxConsecutiveOnes(nums))

}
