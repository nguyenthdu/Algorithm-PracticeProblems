package main

import "fmt"

func Selection(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}

		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	Selection(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
