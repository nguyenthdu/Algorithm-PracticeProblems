package main

import "fmt"

func BubleSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	BubleSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
