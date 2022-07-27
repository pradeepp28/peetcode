package main

import "fmt"

func bitonicPoint(nums []int) int {
	n := len(nums)
	l, h := 0, n-1
	for l <= h {
		mid := l + (h-l)/2
		if mid < n-1 && nums[mid] > nums[mid+1] && mid > 0 && nums[mid] > nums[mid-1] {
			return nums[mid]
		} else if mid == 0 || nums[mid] < nums[mid-1] {
			h = mid - 1
		} else if mid == n-1 || nums[mid] < nums[mid+1] {
			l = mid + 1
		}
	}
	if l == n {
		return nums[n-1]
	}
	if h == -1 {
		return nums[0]
	}
	return -1
}
func main() {
	fmt.Println(bitonicPoint([]int{1, 15, 25, 45, 42, 21, 17, 12, 11})) // 45
	fmt.Println(bitonicPoint([]int{1, 45, 47, 50, 5}))                  // 50
	fmt.Println(bitonicPoint([]int{1, 2, 3, 4, 5, 6, 7}))               // 7
	fmt.Println(bitonicPoint([]int{7, 6, 5, 4, 3, 2, 1}))               // 7
}
