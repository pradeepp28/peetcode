package main

import "fmt"

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
func merge(first, second []int) []int {
	res := make([]int, len(first)+len(second))
	i, j, k := 0, 0, 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			res[k] = first[i]
			i = i + 1
		} else {
			res[k] = second[j]
			j = j + 1
		}
		k = k + 1
	}
	for ; i < len(first); i++ {
		res[k] = first[i]
		k = k + 1
	}
	for ; j < len(second); j++ {
		res[k] = second[j]
		k = k + 1
	}
	return res
}

func mergeSort(l, h int, nums []int) []int {
	if l == h {
		return nums[l : l+1]
	}
	for l < h {
		mid := l + ((h - l) / 2)
		first := mergeSort(l, mid, nums)
		second := mergeSort(mid+1, h, nums)
		return merge(first, second)
	}
	return []int{}
}

func threeSum(nums []int) [][]int {
	var res [][]int
	nums = mergeSort(0, len(nums)-1, nums)
	fmt.Println(nums, len(nums))
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j = j + 1
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j = j + 1
				// k = k - 1
			} else if sum < 0 {
				j = j + 1
			} else {
				k = k - 1
			}
		}
	}
	return res
}
