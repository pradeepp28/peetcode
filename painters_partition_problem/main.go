package main

import "fmt"

func isAllocationPossible(boards []int, painters int, maxTime int) bool {
	count := 1
	n := len(boards)
	timeTaken := 0
	for i := 0; i < n; i++ {
		if timeTaken+boards[i] <= maxTime {
			timeTaken += boards[i]
		} else {
			timeTaken = boards[i]
			count += 1
			if count > painters {
				return false
			}
		}
	}
	return true
}
func minTime(boards []int, painters int) int {
	n := len(boards)
	maxLength := boards[(n-1)/2]
	totalLength := 0
	for i := 0; i < n; i++ {
		totalLength += boards[i]
	}
	l := maxLength
	h := totalLength
	res := 0
	for l <= h {
		mid := l + (h-l)/2
		if isAllocationPossible(boards, painters, mid) {
			res = mid
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res

}
func main() {
	fmt.Println(minTime([]int{5, 10, 30, 20, 15}, 3)) //35
}
