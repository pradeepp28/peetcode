package main

import "fmt"

func helper(nums []int, n int, i int, curr []int, res *[][]int) {
	if i == n {
		*res = append(*res, curr)
		return
	}
	helper(nums, n, i+1, append(curr, nums[i]), res)
	helper(nums, n, i+1, curr, res)
}
func subsets(nums []int) [][]int {
	var res [][]int
	helper(nums, len(nums), 0, []int{}, &res)
	return res
}
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
