package main

func backtrack1(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(s) != 0 && (s[0] == p[0] || p[0] == '?') {
		return backtrack1(s[1:], p[1:])
	}

	if p[0] == '*' {
		return (len(s) > 0 && backtrack1(s[1:], p)) || backtrack1(s, p[1:])
	}
	return false
}
