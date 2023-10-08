package main

import "fmt"

func Insertion(arr []int) {
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--

		}
		arr[j+1] = key

	}
}
func main() {
	arr := []int{7, 9, 2, 5, 3, 1}
	Insertion(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
