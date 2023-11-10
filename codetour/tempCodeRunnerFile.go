func findMaxSum(N [][]int) int {
	size := len(N)
	sum := 0
	for i := 0; i < size; i++ {
		sum += N[i][i]
	}
	return sum % 2027
}