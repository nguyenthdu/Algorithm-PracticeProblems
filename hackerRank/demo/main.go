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
 * Complete the 'reverseArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func reverseArray(a []int32) []int32 {
	b := make([]int32, len(a))
	// Write your code here
	for i := 0; i < len(a); i++ {
		b[i] = a[len(a)-1-i]
	}
	return b

}
func dynamicArray(n int32, queries [][]int32) []int32 {
	// Write your code here
	var lastAnswer int32 = 0
	var result []int32
	var seqList [][]int32
	for i := 0; i < int(n); i++ {
		seqList = append(seqList, []int32{})
	}
	for _, query := range queries {
		x := query[1]
		y := query[2]
		index := (x ^ lastAnswer) % n
		if query[0] == 1 {
			seqList[index] = append(seqList[index], y)
		} else {
			lastAnswer = seqList[index][y%int32(len(seqList[index]))]
			result = append(result, lastAnswer)
		}
	}
	return result
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)

	}

	res := reverseArray(arr)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
