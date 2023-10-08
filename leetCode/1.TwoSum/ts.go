package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // bao gom key la so va value la vi tri cua so do
	
	for i := 0; i < len(nums); i++ {
		key := target - nums[i] //lay phan bu vi du target =7 i =2 key =5
		if m[key] > 0 {// neu key >0 thi tra ve vi tri cua key va vi tri hien tai

			return []int{m[key] - 1, i}
		}
		m[nums[i]] = i + 1
	}
	return nil
}
func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if val, ok := m[complement]; ok { // neu ok thi tra ve true va val la gia tri cua key
			//ok là biến kiểu bool, nếu key có tồn tại trong map thì ok = true, ngược lại ok = false
			return []int{val, i} // tra ve vi tri cua so do va vi tri hien tai
		}

		m[num] = i// neu khong thi gan key la so va value la vi tri cua so do
	}
	return []int{}
}
