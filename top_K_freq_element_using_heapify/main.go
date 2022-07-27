package main

import "fmt"

type freq struct {
	num   int
	count int
}

func heapify(freqList []freq, n, pos int) {
	if len(freqList) == 0 {
		return
	}
	largest := pos
	left := 2*pos + 1
	right := 2*pos + 2
	if left < n && freqList[left].count > freqList[largest].count {
		largest = left
	}
	if right < n && freqList[right].count > freqList[largest].count {
		largest = right
	}
	if largest != pos {
		freqList[pos], freqList[largest] = freqList[largest], freqList[pos]
		heapify(freqList, n, largest)
	}
}

func heapSort(freqList []freq) {
	n := len(freqList)
	for i := (n - 1) / 2; i >= 0; i-- {
		heapify(freqList, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		freqList[0], freqList[i] = freqList[i], freqList[0]
		heapify(freqList, i, 0)
	}
}

func popFromHeap(freqList []freq) []freq {
	if len(freqList) == 0 {
		return []freq{}
	}
	n := len(freqList)
	freqList[0], freqList[n-1] = freqList[n-1], freqList[0]
	heapify(freqList, len(freqList)-1, 0)
	return freqList[0 : n-1]
}

func topKFrequent1(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}

	freqList := make([]freq, len(m))
	i := 0
	for k, v := range m {
		freqList[i].num = k
		freqList[i].count = v
		i++
	}
	fmt.Printf("%+v\n", freqList)

	heapSort(freqList)
	fmt.Printf("%+v\n", freqList)
	res := make([]int, k)
	n := len(freqList)
	for i := 0; i < k; i++ {
		res[i] = freqList[n-i-1].num
	}
	return res
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}

	freqList := make([]freq, len(m))
	i := 0
	for k, v := range m {
		freqList[i].num = k
		freqList[i].count = v
		i++
	}
	fmt.Printf("%+v\n", freqList)

	n := len(freqList)
	for i := (n - 1) / 2; i >= 0; i-- {
		heapify(freqList, n, i)
	}
	fmt.Printf("%+v\n", freqList)
	res := make([]int, k)
	// n := len(freqList)
	for i := 0; i < k; i++ {
		res[i] = freqList[0].num
		freqList = popFromHeap(freqList)
	}
	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{3, 0, 1, 0}, 1))
	fmt.Println(topKFrequent([]int{-1, 1, 4, -4, 3, 5, 4, -2, 3, -1}, 3))
}
