package main

import "fmt"

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here

	var result int32 = 0
	for i := 0; i < len(s); i++ {
		var count int32 = 0
		var sum int32

		for j := i; j < len(s); j++ {
			sum += s[j]
			count++
			if sum == d {
				if count == m {
					result++
				}
				break
			}

		}

	}
	return result

}

func main() {
	s := []int32{1, 1, 1, 1, 1, 1}
	result := birthday(s, 3, 2)
	fmt.Println(result)
}
