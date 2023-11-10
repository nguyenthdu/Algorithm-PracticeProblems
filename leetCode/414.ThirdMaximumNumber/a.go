package main

import "sort"

func thirdMax(nums []int) int {

	sort.Ints(nums) //tại sao lại sort: để tìm ra số lớn thứ 3

	unique := make(map[int]bool)
	result := []int{}
	for _, num := range nums {
		if !unique[num] { //neu chua xuat hien
			result = append(result, num)
			unique[num] = true
		}
	}
	if len(result) < 3 {
		return result[len(result)-1]
	}
	return result[len(result)-3]

}
