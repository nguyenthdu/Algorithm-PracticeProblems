package main

import "fmt"

func replaceElements(arr []int) []int {
	var max int = -1
	for i := len(arr) - 1; i >= 0; i-- { //chạy từ cuối mảng về đầu mảng
		temp := arr[i]  //lưu lại giá trị của arr[i] trước khi gán max cho arr[i]
		arr[i] = max    //gán max cho arr[i]
		if temp > max { //nếu giá trị của arr[i] lớn hơn max thì gán max = arr[i]
			max = temp //gán max = arr[i]
		}
	}
	return arr
}
func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))

}
