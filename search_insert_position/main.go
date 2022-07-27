package main

import "fmt"

func search(nums []int, target int) int {
	l, h := 0, len(nums)-1
	mid := 0
	for l <= h {
		mid = l + ((h - l) / 2)
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	if target < nums[mid] {
		return mid
	}
	return mid + 1
}
func main() {
	fmt.Println(search([]int{1, 3, 5, 6}, 5)) //5
	fmt.Println(search([]int{1, 3, 5, 6}, 2)) //1
	fmt.Println(search([]int{1, 3, 5, 6}, 7)) //4
}
