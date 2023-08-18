package main

import "fmt"

func numberOfSteps(num int) int {
	var count int
	for {
		if num == 0 {
			break
		} else if num%2 == 0 {
			num /= 2
			count++
		} else {
			num -= 1
			count++
		}
	}
	return count

}

func main() {
	fmt.Println(numberOfSteps(14))

}
