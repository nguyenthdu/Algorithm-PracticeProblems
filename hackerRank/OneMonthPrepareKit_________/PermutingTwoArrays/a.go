package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'twoArrays' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY A
 *  3. INTEGER_ARRAY B
 */

func twoArrays(k int32, A []int32, B []int32) string {
	//permute A
	n := len(A)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if A[i] > A[j] {
				A[i], A[j] = A[j], A[i]
			}
			if B[i] < B[j] {
				B[i], B[j] = B[j], B[i]
			}
		}
	}
	for i := 0; i < n; i++ {
		if A[i] > A[i+1] {
			A[i], A[i+1] = A[i+1], A[i]
		}
		if B[i] < B[i+1] {
			B[i], B[i+1] = B[i+1], B[i]
		}
	}
	//check
	for i := 0; i < n; i++ {
		if A[i]+B[i] < k {
			return "NO"
		}
	}
	return "YES"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		k := int32(kTemp)

		ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var A []int32

		for i := 0; i < int(n); i++ {
			AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			A = append(A, AItem)
		}

		BTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var B []int32

		for i := 0; i < int(n); i++ {
			BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
			checkError(err)
			BItem := int32(BItemTemp)
			B = append(B, BItem)
		}

		result := twoArrays(k, A, B)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
