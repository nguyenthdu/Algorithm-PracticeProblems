package main

import (
	"fmt"
)

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 */

func matchingStrings(strings []string, queries []string) []int32 {
	// Write your code here
	var result = []int32{}
	for q := 0; q < len(queries); q++ {
		count := 0

		for s := 0; s < len(strings); s++ {
			if queries[q] == strings[s] {
				count++
			}
		}
		result = append(result, int32(count))

	}
	return result
}
func matchingStrings1(strings []string, queries []string) []int32 {
	// Write your code here
	var result = []int32{}
	m := make(map[string]int)
	for _, s := range strings {
		m[s]++ //giai thich: neu s da co trong map thi tang len 1, neu chua co thi gan bang 1
		//key la string, value la so lan xuat hien cua string do
	}
	for _, q := range queries {
		result = append(result, int32(m[q]))

	}

	return result

}

func main() {
	s := []string{"def", "de", "fgh", "xzxb"}
	q := []string{"de", "lmn", "fgh"}
	fmt.Println(matchingStrings1(s, q))
	// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// 	checkError(err)

	// 	defer stdout.Close()

	// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// 	stringsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// 	checkError(err)

	// 	var strings []string

	// 	for i := 0; i < int(stringsCount); i++ {
	// 		stringsItem := readLine(reader)
	// 		strings = append(strings, stringsItem)
	// 	}

	// 	// queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// 	checkError(err)

	// 	var queries []string

	// 	for i := 0; i < int(stringsCount); i++ {
	// 		queriesItem := readLine(reader)
	// 		queries = append(queries, queriesItem)
	// 	}

	// 	res := matchingStrings(strings, queries)

	// 	for i, resItem := range res {
	// 		fmt.Fprintf(writer, "%d", resItem)

	// 		if i != len(res)-1 {
	// 			fmt.Fprintf(writer, "\n")
	// 		}
	// 	}

	// 	fmt.Fprintf(writer, "\n")

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
}
