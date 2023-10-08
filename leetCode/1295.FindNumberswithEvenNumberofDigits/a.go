package main

import (
	"math"
	"strconv"
)

func findNumbers(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		n := countN(nums[i])
		if n%2 == 0 {
			result++
		}
	}
	return result
}
func countN(number int) int {
	var count int
	for number > 0 {
		number /= 10
		count++
	}
	return count
}
func findNumbers1(nums []int) int {
	var count int
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			count++
		}
	}
	return count
}

func findNumbers2(nums []int) int {
	var count int
	for _, num := range nums {
		if num > 9 && num < 100 || num > 999 && num < 10000 || num == 100000 {
			count++
		}
	}
	return count
	//giải thích: số có 2 chữ số hoặc 4 chữ số hoặc 6 chữ số

}

// using logarit
func findNumbers3(nums []int) int {
	var count int
	for _, num := range nums {
		if int(math.Log10(float64(num)))%2 == 1 {
			count++
		}
	}
	return count
}
