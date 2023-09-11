package main

import "fmt"

// đếm tổng các chữ số là chẵn hay lẻ
func findNumbers(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		n := 0
		for nums[i] > 0 {
			nums[i] = nums[i] / 10
			n++
		}
		if n%2 == 0 {
			count++
		}
	}

	return count

}
func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))

}
