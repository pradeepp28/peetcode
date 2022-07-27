package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var res int
	for i, j := 0, 0; j < len(s); j++ {
		if _, ok := m[s[j]]; ok {
			i = max(i, m[s[j]])
		}
		res = max(res, j-i+1)
		m[s[j]] = j + 1
	}
	return res
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))      // 3
	fmt.Println(lengthOfLongestSubstring("abcdeafbdgcbb")) // 7
	fmt.Println(lengthOfLongestSubstring("bbbb"))          // 1
}
