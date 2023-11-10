package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	listCus := make([]int, 0)
	for _, i := range accounts {
		var amount int
		for _, j := range i {
			amount += j
		}
		listCus = append(listCus, amount)
	}
	//-> độ phức tạp O(n^2)
	fmt.Println(listCus)

	maxCus := listCus[0]
	for _, val := range listCus {
		if val > maxCus {
			maxCus = val
		}

	}
	return maxCus

}

//TODO: Su dung thuat toan tinh kiem tuyen tinh
func maximumWealth2(accounts [][]int) int {
	maxCus := 0
	for _, i := range accounts { // duyet qua tat ca cac tai khoan
		sum := 0
		for _, j := range i { // tinh tong tien cua tung tai khoan
			sum += j
		}
		if sum > maxCus { // so sanh
			maxCus = sum
		}
	}
	return maxCus
}

func main() {
	account := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	fmt.Println(maximumWealth(account))
}
