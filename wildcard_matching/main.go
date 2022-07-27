package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "*"))          // true
	fmt.Println(backtrack1("aa", "*"))       // true
	fmt.Println(backtrack1("aab", "*b"))     // true
	fmt.Println(backtrack1("aab", "c*a*b"))  // false
	fmt.Println(backtrack1("ab", "?b"))      // true
	fmt.Println(backtrack1("ab", "?c"))      // false
	fmt.Println(backtrack1("adceb", "*a*b")) // true

}

var (
	m [][]*bool
)

func backtrack(i, j int, s, p string) bool {
	if m[i][j] != nil {
		return *m[i][j]
	}
	if s[i-1] == p[j-1] || p[j-1] == '?' {
		return backtrack(i-1, j-1, s, p)
	}
	if p[j-1] == '*' {
		return backtrack(i, j-1, s, p) || backtrack(i-1, j, s, p)
	}
	return false
}

func isMatch(s string, p string) bool {

	m = make([][]*bool, len(s)+1)
	for i := range m {
		m[i] = make([]*bool, len(p)+1)
	}
	t := true
	m[0][0] = &t

	for i := 1; i < len(s)+1; i++ {
		f := false
		m[i][0] = &f
	}
	for j := 1; j < len(p)+1; j++ {
		if p[j-1] == '*' {
			m[0][j] = m[0][j-1]
		} else {
			f := false
			m[0][j] = &f
		}
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			b := backtrack(i, j, s, p)
			m[i][j] = &b
		}
	}

	for i := 0; i < len(s)+1; i++ {
		for j := 0; j < len(p)+1; j++ {
			// fmt.Printf("%t ", *m[i][j])
		}
		// fmt.Println("")
	}
	return *m[len(s)][len(p)]
}
