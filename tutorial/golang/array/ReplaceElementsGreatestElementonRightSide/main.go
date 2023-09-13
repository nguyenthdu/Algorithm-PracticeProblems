package main

import "fmt"

func replaceElements(arr []int) []int {
	var max int = -1
	for i := len(arr) - 1; i >= 0; i-- {
		temp := arr[i]
		arr[i] = max
		if temp > max {
			max = temp
		}
	}
	return arr
}
func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))

}
