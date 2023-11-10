package main

// xoa phan tu trung nhau trong mang
func removeDuplicates(nums []int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[k] { //neu phan tu tai vi tri i khac phan tu tai vi tri k thi gan phan tu tai vi tri i cho phan tu tai vi tri k va tang
			k++ // k la vi tri cua phan tu khac nhau
			nums[k] = nums[i]
		}
	}
	return k + 1 //giải thích tại sao: vì k là vị trí của phần tử khác nhau nên k+1 là số lượng phần tử khác nhau
}
func removeDuplicates2(nums []int, val int) []int {
	n := len(nums)
	for i := 0; i < n;{
		if val == nums[i] {
			for j := i; j <= n-2; j++ {
				nums[j] = nums[j+1]
			}
			n--
		}else{
			i++
		}
	}
	return nums
}
func main() {
	nums := []int{1, 2, 3, 3, 6}
	removeDuplicates2(nums, 3)

}
