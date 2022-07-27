package main

import "fmt"

func helper(nums []int, target, l, h int) int {
	if l > h {
		return -1
	}
	mid := l + (h-l)/2
	if nums[mid] == target {
		return 1
	} else if target < nums[mid] {
		return helper(nums, target, l, mid-1)
	} else {
		return helper(nums, target, mid+1, h)
	}
}
func search(nums []int, target int) int {
	return helper(nums, target, 0, len(nums)-1)
}
func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 6}, 6)) // 1
	fmt.Println(search([]int{1, 2, 3, 4, 6}, 5)) // -1
}
