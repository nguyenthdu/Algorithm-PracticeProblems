package main

import (
	"sort"
)

func sortedSquares(nums []int) []int {
	result := []int{}
	for _, v := range nums {
		if v < 0 {
			v = v * -1
		}

		result = append(result, v*v)
	}
	sort.Ints(result)
	return result
}
func main() {

}
