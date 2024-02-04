package main

import "fmt"

func BSearch_Recursion(list []int, key int) int {
	left := 0
	right := len(list) - 1
	for left <= right {
		mid := (left + right) / 2
		if key == list[mid] {
			return mid
		} else if list[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	list := []int{1, 2, 4, 5, 6}
	fmt.Println(BSearch_Recursion(list, 5))
}
