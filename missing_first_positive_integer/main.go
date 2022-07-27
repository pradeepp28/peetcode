package main

import "fmt"

type Stack struct {
	elements []byte
}

func (s *Stack) Push(e byte) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() (byte, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	l := len(s.elements)
	top := s.elements[l-1]
	s.elements = s.elements[:l-1]
	return top, true
}

var (
	m = map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
)

func isValid(s string) bool {
	stack := &Stack{}
	for i := 0; i < len(s); i++ {
		switch v := s[i]; v {
		case '(', '{', '[':
			stack.Push(v)
		default:
			if element, ok := stack.Pop(); !ok || element != m[v] {
				return false
			}
		}
	}
	_, exists := stack.Pop()
	if exists {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))   // true
	fmt.Println(isValid("({"))   // false
	fmt.Println(isValid("({})")) // true
	fmt.Println(isValid("({)}")) // false
}
