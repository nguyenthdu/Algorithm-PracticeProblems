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
func Selection2(a []int) {
	for i := 0; i < len(a); i++ {
		minIndex := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}

		}
		if minIndex != i {
			temp := a[minIndex]
			a[minIndex] = a[i]
			a[i] = temp
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
