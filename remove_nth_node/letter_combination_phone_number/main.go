package main

func main() {
	letterCombinations("2")
}

var res []string
var keypadMapping = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func backtrack(i int, digits string, curString []byte) {
	if len(digits) == len(curString) {
		res = append(res, string(curString))
		return
	}
	for _, c := range []byte(keypadMapping[digits[i]]) {
		backtrack(i+1, digits, append(curString, c))
	}
}
func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	res = []string{}
	backtrack(0, digits, []byte(""))
	return res
}
