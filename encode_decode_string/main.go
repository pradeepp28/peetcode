package main

import (
	"fmt"
	"strconv"
)

func Encode(input []string) string {
	var res string
	for _, str := range input {
		res += strconv.Itoa(len(str)) + "#" + str
	}
	return res
}
func Decode(str string) []string {
	var res []string
	i := 0
	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}
		l := str[i:j]
		n, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("Failed to convert, %v", err)
		}
		res = append(res, str[j+1:j+1+n])
		i = j + 1 + n
	}
	return res
}
func main() {
	e := Encode([]string{"neet", "cod#e"})
	fmt.Println(e)
	d := Decode(e)
	fmt.Println(d)
}
