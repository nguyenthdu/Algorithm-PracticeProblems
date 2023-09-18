package main

import "fmt"

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	if arr[0] >= arr[1] {
		return false
	}
	peak := false
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}
		if !peak && arr[i] < arr[i-1] {//3 2 1
			peak = true
		}
		if peak && arr[i] > arr[i-1] {//1 2 3
			return false
		}
	}
	return peak
}
func main() {
	arr := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(arr))
}
