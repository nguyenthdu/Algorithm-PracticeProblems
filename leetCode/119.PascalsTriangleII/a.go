package main

func getRow(rowIndex int) []int {

	result := []int{1}
	for i := 1; i <= rowIndex; i++ {
		result = append(result, result[i-1]*(rowIndex-i+1)/i)
		// giải thích công thức:
		// result[i-1] là phần tử trước đó
		// (rowIndex-i+1)/i là tỉ lệ của phần tử hiện tại so với phần tử trước đó
		// ví dụ: [1, 3, 3, 1] thì 3/1 = 3, 3/2 = 1.5, 1/3 = 0.3333
		// nếu như nhân tỉ lệ đó với phần tử trước đó thì sẽ ra phần tử hiện tại
		// ví dụ: 3 * 3 = 9, 3 * 1.5 = 4.5, 1 * 0.3333 = 0.3333
		// như vậy ta sẽ có được dãy [1, 3, 3, 1]

	}


	return result
}
