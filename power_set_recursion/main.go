package main

import (
	"fmt"
	"sort"
)

func powerSet(s string, n int, idx int, curr string, res *[]string) {
	if idx == n {
		*res = append(*res, curr)
		return
	}
	// include
	powerSet(s, n, idx+1, curr+string(s[idx]), res)
	// exclude
	powerSet(s, n, idx+1, curr, res)
}
func subsequences(s string) []string {
	res := &[]string{}
	powerSet(s, len(s), 0, "", res)
	fmt.Printf("%d - %+v\n", len(*res), *res)
	sort.Strings(*res)
	fmt.Printf("%d - %+v\n", len(*res), *res)
	return (*res)[1:]
}
func main() {
	fmt.Println(subsequences("ab"))
	fmt.Println(subsequences("abc"))
}
