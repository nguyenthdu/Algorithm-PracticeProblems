package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(s string) string {
	// Write your code here
	var check [26]int //mang 26 phan tu vi tieng anh co 26 chu cai
	for _, c := range s {
		if c >= 'a' && c <= 'z' { //neu la chu thuong thi tang check len 1
			check[c-'a']++
		} else if c >= 'A' && c <= 'Z' { //neu la chu hoa thi tang check len 1
			check[c-'A']++
		}
	}
	for _, v := range check {
		if v == 0 {
			return "not pangram"
		}
	}
	return "pangram"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := pangrams(s)

	fmt.Fprintf(writer, "%s\n", result)

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
