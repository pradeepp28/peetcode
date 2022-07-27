package main

import "fmt"
var m map[int]int
func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
func dp(curr int, n int, cost []int) int {
    if v, ok := m[curr]; ok {
        return v
    }
    if curr >= n {
        return 0
    }
    m[curr] = cost[curr] + min(dp(curr+1, n, cost), dp(curr+2, n, cost))
    return m[curr]
}
func minCostClimbingStairs(cost []int) int {
    m = make(map[int]int)
    return min(dp(0, len(cost), cost), dp(1, len(cost), cost))
}
func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))                         //15
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) //6
}
