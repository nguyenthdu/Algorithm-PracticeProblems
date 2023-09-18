package main

import "fmt"

func removeDuplicates(arr []int) []int {
	unique := make(map[int]bool)
	result := []int{}
	for _, num := range arr {
		if !unique[num] { //neu chua xuat hien
			result = append(result, num)
			unique[num] = true
		}
	}
	return result

}
func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(arr)
	fmt.Println(result)

}
