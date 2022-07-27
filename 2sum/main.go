package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		r := target - numbers[i]
		if v, ok := m[numbers[i]]; ok {
			return []int{v, i + 1}
		}
		m[r] = i + 1
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
