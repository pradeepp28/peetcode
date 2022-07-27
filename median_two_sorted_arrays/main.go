package main

import (
	"fmt"
)

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
func findMedian(a []int64, b []int64) int64 {
	m := len(a)
	n := len(b)
	total := m + n
	half := total / 2
	posInf := max(a[m-1], b[n-1]) + 10 // positive infinity
	negInf := min(a[0], b[0]) - 10     // negative infinity

	if n < m {
		a, b = b, a
		// m, n = n, m
	}

	l, r := 0, len(a)-1
	i := (l + r) / 2
	j := half - i - 2
	for {
		var aleft, aright, bleft, bright int64
		if i >= 0 {
			aleft = a[i]
		} else {
			aleft = negInf
		}
		if (i + 1) < len(a) {
			aright = a[i+1]
		} else {
			aright = posInf
		}
		if j >= 0 {
			bleft = b[j]
		} else {
			bleft = negInf
		}
		if (j + 1) < len(b) {
			bright = b[j+1]
		} else {
			bright = posInf
		}
		if aleft <= bright && bleft <= aright {
			if total%2 == 1 {
				return min(aright, bright)
			} else {
				return (max(aleft, bleft) + min(aright, bright)) / 2
			}
		} else if aleft > bright {
			i = i - 1
			j = j + 1
		} else {
			j = j - 1
			i = i + 1
		}
	}
}
func main() {
	fmt.Println(findMedian([]int64{2, 3, 6, 7}, []int64{1, 5, 8, 10, 18, 20}))
	fmt.Println(findMedian([]int64{1, 3}, []int64{2}))
	fmt.Println(findMedian([]int64{1, 2}, []int64{3, 4}))
	fmt.Println(findMedian([]int64{1, 2, 3, 4, 5, 6, 7, 8}, []int64{1, 2, 3, 4, 5}))
}
