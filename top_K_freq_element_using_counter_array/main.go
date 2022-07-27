package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}

	count := make([][]int, len(nums)+1)
	for k, v := range m {
		count[v] = append(count[v], k)
	}
	var res []int
	for i := len(count) - 1; i > 0 && len(res) < k; i-- {
		for j := 0; j < len(count[i]) && len(res) < k; j++ {
			res = append(res, count[i][j])
		}
	}
	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))                 // 1,2
	fmt.Println(topKFrequent([]int{3, 0, 1, 0}, 1))                       // 0
	fmt.Println(topKFrequent([]int{-1, 1, 4, -4, 3, 5, 4, -2, 3, -1}, 3)) // -1, 4, 3
	fmt.Println(topKFrequent([]int{1, 1, 2, 2, 3, 3}, 2))
}
