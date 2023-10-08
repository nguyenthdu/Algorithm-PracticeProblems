package main

import "fmt"

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == 2*arr[j] {
				return true
			}
		}
	}
	return false
}
func main() {
	arr := []int{10, 2, 5, 3}
	fmt.Println(checkIfExist(arr))

}
