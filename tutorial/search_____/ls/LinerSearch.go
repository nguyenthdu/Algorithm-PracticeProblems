package main

import "fmt"

func LSearch(arr []int, n int, key int) {
	flag := 0

	for i := 0; i < n; i++ {
		if arr[i] == key {
			fmt.Printf("found at position %d", i)
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("not found")
	}

}
func LSearch1(arr []int, n int, key int) int {
	find := -1

	for i := 0; i < n; i++ {
		if arr[i] == key {
			find = i
			break
		}
	}
	return find
}
func main() {

	n := 5
	k := 5
	arr := []int{2, 4, 5, 6, 7}
	LSearch(arr, n, k)
}
