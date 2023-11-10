package main

func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ { // tại vị trí i thì ta có 2 cách để đến đó là từ i-1 hoặc i-2
		cost[i] += min(cost[i-1], cost[i-2]) // ta chọn cách nào có cost nhỏ nhất
	}
	return min(cost[len(cost)-1], cost[len(cost)-2]) // tại vị trí cuối cùng thì ta có 2 cách để đến đó là từ i-1 hoặc i-2

}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCostClimbingStairs1(cost []int) int {
	first := cost[0]
	second := cost[1]

	for i := 2; i < len(cost); i++ {
		if first > second {
			first, second = second, second+cost[i] //
		} else {
			first, second = second, first+cost[i]
		}
	}

	if first < second {
		return first
	}
	return second
}
func minCostClimbingStairs2(cost []int) int {
	first := 0
	second := 0

	for i := 0; i < len(cost); i++ {
		step := cost[i] + min(first, second)
		first = second
		second = step

	}

	return min(first, second)
}

func main() {
	//[1,100,1,1,1,100,1,1,100,1]
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	minCostClimbingStairs(cost)

}
