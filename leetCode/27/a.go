package main

func removeDuplicates2(nums []int, val int) []int {
	n := len(nums)
	for i := 0; i < n; {
		if val == nums[i] {
			for j := i; j <= n-2; j++ {
				nums[j] = nums[j+1]
			}
			n--
		} else {
			i++
		}
		
	}
	return nums
}
func main() {
	nums := []int{1, 2, 3, 3, 6}
	removeDuplicates2(nums, 3)

}
