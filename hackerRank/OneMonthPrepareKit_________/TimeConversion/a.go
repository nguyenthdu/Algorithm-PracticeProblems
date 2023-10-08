package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	t, _ := time.Parse("03:04:05PM", s) //giai thich: 03:04:05PM là format của s. s là string đầu vào
	return t.Format("15:04:05")         //giai thich: 15:04:05 là format của t. t là time sau khi được parse
	// 15:04:05 là format mac dinh của time
	// 03:04:05PM là format mac dinh của string

}

// khong dung time package
func timeConversion1(s string) string {
	// Write your code here
	// 07:05:45PM
	// 0123456789
	// 19:05:45
	// 01234567
	//PM
	if s[8] == 'P' {
		if s[0:2] == "12" { // 12:05:45PM->12:05:45
			return s[0:8]
		}
		return strconv.Itoa(12+int(s[1]-'0')) + s[2:8] // 07:05:45PM->19:05:45
		// 12+int(s[1]-'0') = 12+int('7'-'0') = 12+7 = 19
		//strconv.Itoa() chuyen int ve string
	}
	//AM
	if s[0:2] == "12" {
		return "00" + s[2:8]
	}

	return s[0:8]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
