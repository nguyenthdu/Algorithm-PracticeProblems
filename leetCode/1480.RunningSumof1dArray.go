package main

import "fmt"
// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
// Return the running sum of nums.
//Tong cua cac phan tu phia truoc

func runningSum(nums []int) []int {
	n := len(nums)
	arr := make([]int, n)
	arr[0] = nums[0]//bat dau 2 phan tu dau tien bang nhau

	for i := 1; i < len(nums); i++ {
		arr[i] = nums[i] + arr[i-1]
	}
	return arr

}

func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func main() {
	a := []int{1, 2, 3, 5}
	fmt.Println(runningSum2(a))
}