package main

import (
	"fmt"
)

func isAllocationPossible(packets []int, channels int, maxTime int) ([][]int, bool) {
	count := 1
	n := len(packets)
	timeTaken := 0
	res := make([][]int, channels)
	for i := 0; i < n; i++ {
		if timeTaken+packets[i] <= maxTime {
			timeTaken += packets[i]
		} else {
			timeTaken = packets[i]
			count += 1
			if count > channels {
				return nil, false
			}
		}
		res[count-1] = append(res[count-1], packets[i])
	}
	return res, true
}
func medianOfChannels(packets []int, channels int) int {
	// sort.Ints(packets)
	n := len(packets)
	maxPackets := packets[(n-1)/2]
	totalPackets := 0
	for i := 0; i < n; i++ {
		totalPackets += packets[i]
	}
	l := maxPackets
	h := totalPackets
	var res [][]int

	for l <= h {
		mid := l + (h-l)/2
		var ok bool
		res, ok = isAllocationPossible(packets, channels, mid)
		if ok {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	fmt.Println(res)
	var median int
	for i := range res {
		l := len(res[i])
		currMedian := 0
		if l%2 == 0 {
			sum := res[i][(l-1)/2] + res[i][(l-1)/2+1]
			if sum%2 != 0 {
				currMedian = sum/2 + 1
			} else {
				currMedian = sum / 2
			}
		} else {
			currMedian = res[i][(l-1)/2]
		}
		median += currMedian
	}
	return median

}
func main() {
	fmt.Println(medianOfChannels([]int{1, 2, 2, 5, 3}, 2))
	fmt.Println(medianOfChannels([]int{1, 2, 3}, 2))
}
