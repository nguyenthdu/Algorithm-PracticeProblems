package main

import "fmt"

//TODO: Nhập vào 1 chuỗi lưu vào danh sách, nhập vào nội dung cần tìm
// In ra số lần xuất hiện
// In ra vị trí xuất hiện
// In ra độ dài nội dung cần tìm

func Stringhandling(str, text string) {
	n := 0
	strtt := len(str)
	texttt := len(text)
	for i := 0; i < strtt-texttt; i++ {
		if str[i:i+texttt] == text {
			n++
			fmt.Println("Vi tri xuat hien la: ", i)
		}

	}
	fmt.Println("So lan xuat hien cua la: ", n)

}

func main() {
	var str string = "ab ab ab cd cd "
	var text string = "ab"
	// fmt.Println("Enter string: ")
	// fmt.Scan(str)
	// fmt.Println("Enter the text you want to find: ")
	// fmt.Scan(text)
	fmt.Println("Dai cua chuoi can tim la: ", len(text))
	Stringhandling(str, text)
}
