package main

import "fmt"

func printPattern(nums []int, n int) []int {
	if n <= 0 {
		fmt.Print(n)

		return append(nums, n)
	}
	fmt.Print(n, " ")
	nums = printPattern(append(nums, n), n-5)
	fmt.Print(" ", n)
	return append(nums, n)
}
func main() {
	nums := printPattern([]int{}, 10)
	fmt.Println()
	fmt.Println(nums)
	// printPattern(16)
	// fmt.Println()
}
