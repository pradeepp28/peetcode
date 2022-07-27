package main

import "fmt"

func reverseWord(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(append([]byte(reverseWord(s[1:])), s[0]))

}
func main() {
	fmt.Println(reverseWord("Geeks"))
}
