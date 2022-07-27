package main

func minNumberOfPages(pages []int, m int) int {
	totalPages := 0
	n := len(pages)
	for i := 0; i < n; i++ {
		totalPages = totalPages + pages[i]
	}
	for {
		count := 0
		minPages := totalPages / 2
		allocate := 0
		for i := 0; i < n; i++ {
			allocate = allocate + pages[i]
			if allocate > minPages {
				allocate = pages[i]
				count = count + 1
			}
		}
		if count == m {
			return totalPages - allocate
		}
	}

}
func main() {

}
