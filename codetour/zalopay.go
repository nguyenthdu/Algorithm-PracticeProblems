package main

import "fmt"

//cho ma trận N * N
//Chọn ra N  ô vuông sao cho không có 2 ô được chọn nào cùng dòng hoặc cùng cột. Tìm tổng lớn nhất có thể của N ô vuông chia dư cho 2027
//Input
//Dòng đầu tiên là số nguyên dương N (1 ≤ N ≤ 1000)

// type Cell struct {
// 	Value, Row, Col int
// }

// type ByValue []Cell

// func (a ByValue) Len() int           { return len(a) }
// func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByValue) Less(i, j int) bool { return a[i].Value > a[j].Value }

// func findMaxSum(N [][]int) int {
// 	var cells []Cell
// 	size := len(N)

// 	// Tạo danh sách các ô với giá trị và vị trí
// 	for i := 0; i < size; i++ {
// 		for j := 0; j < size; j++ {
// 			cells = append(cells, Cell{Value: N[i][j], Row: i, Col: j})
// 		}
// 	}

// 	// Sắp xếp các ô theo giá trị giảm dần
// 	sort.Sort(ByValue(cells))

// 	// Chọn các ô không có 2 ô cùng hàng hoặc cùng cột
// 	var selectedCells []Cell
// 	selectedRows := make(map[int]bool)
// 	selectedCols := make(map[int]bool)

// 	for _, cell := range cells {
// 		if !selectedRows[cell.Row] && !selectedCols[cell.Col] {
// 			selectedCells = append(selectedCells, cell)
// 			selectedRows[cell.Row] = true
// 			selectedCols[cell.Col] = true
// 		}
// 	}

// 	// Tính tổng các ô đã chọn
// 	sum := 0
// 	for _, cell := range selectedCells {
// 		sum += cell.Value
// 	}

// 	return sum % 2027
// }
func findMaxSum(N [][]int) int {
	size := len(N)
	sum := 0
	for i := 0; i < size; i++ {
		sum += N[i][i]
	}
	return sum % 2027
}
func generateMatrix(size int) [][]int {

	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			matrix[i][j] = size*i + j + 1
		}
	}
	return matrix
}

// func findMaxSum1(size int) int {
// 	result := 0
// 	for i := 0; i < size; i++ {
// 		result += i*size + i + 1
// 	}

// 	return result % 2027
// }

func main() {
	var size int
	fmt.Scan(&size)
	matrix := generateMatrix(size)
	fmt.Println(findMaxSum(matrix))
}
