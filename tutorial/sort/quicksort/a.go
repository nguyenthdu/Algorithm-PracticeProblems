package main

import "fmt"

func quicksort(list []int, left int, right int) {
	stack := make([]int, right-left+1)
	top := -1

	top++
	stack[top] = left
	top++
	stack[top] = right

	for top >= 0 {
		right = stack[top]
		top--
		left = stack[top]
		top--

		pivotIndex := partition(list, left, right)

		if pivotIndex-1 > left {
			top++
			stack[top] = left
			top++
			stack[top] = pivotIndex - 1
		}

		if pivotIndex+1 < right {
			top++
			stack[top] = pivotIndex + 1
			top++
			stack[top] = right
		}
	}
}

func partition(list []int, left int, right int) int {
	pivot := list[right]
	i := left - 1

	for j := left; j <= right-1; j++ {
		if list[j] < pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}

	list[i+1], list[right] = list[right], list[i+1]
	return i + 1
}

func main() {
	list := []int{4, 2, 5, 3}
	quicksort(list, 0, len(list)-1)
	fmt.Println(list)
}
