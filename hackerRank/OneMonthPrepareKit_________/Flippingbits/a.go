package main

import "fmt"

/*
 * Complete the 'flippingBits' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER n as parameter.
 */

func main() {
	// Write your code here
	n := 9
	var result int64
	result = 0
	for i := 0; i < 32; i++ {
		if n%2 == 0 {
			result += 1 << i //dich bit sang trai 1 bit, 0011<<1 = 0110, hoặc 0011<<2 = 1100
		}
		n /= 2
	}
	fmt.Println("result: ", result)

}

//using Bitwise NOT
func flippingBits(n int64) int64 {
	return ^n
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// /*
//  * Complete the 'flippingBits' function below.
//  *
//  * The function is expected to return a LONG_INTEGER.
//  * The function accepts LONG_INTEGER n as parameter.
//  */

// func flippingBits(n int64) int64 {
// 	// Write your code here
// 	var result int64
// 	result = 0
// 	for i := 0; i < 32; i++ {
// 		if n%2 == 0 {
// 			result += 1 << i
// 		}
// 		n /= 2
// 	}
// 	return result

// }

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	q := int32(qTemp)

// 	for qItr := 0; qItr < int(q); qItr++ {
// 		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 		checkError(err)

// 		result := flippingBits(n)

// 		fmt.Fprintf(writer, "%d\n", result)
// 	}

// 	writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
