package main

import (
	"fmt"
	"strconv"
)

// func fizzBuzz(n int) []string {
// 	var answer []string
// 	for i := 1; i <= n; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			answer = append(answer, "FizzBuzz")
// 		} else if i%3 == 0 {
// 			answer = append(answer, "Fizz")
// 		} else if i%5 == 0 {
// 			answer = append(answer, "Buzz")
// 		} else {
// 			answer = append(answer, strconv.Itoa(i))
// 		}

//		}
//		return answer
//	}
func fizzBuzz(n int) []string {
	var answer []string
	for i := 1; i <= n; i++ {
		divi3 := i%3 == 0
		divi5 := i%5 == 0
		var strDivi string
		if divi3 {
			strDivi += "Fizz"
		}
		if divi5 {
			strDivi += "Buzz"
		}
		if strDivi == "" {
			strDivi += strconv.Itoa(i)
		}
		answer = append(answer, strDivi)

	}
	return answer
}
func main() {
	fmt.Println(fizzBuzz(3))
}
