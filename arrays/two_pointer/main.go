package two_pointer

import "fmt"

func reverseWords(s string) string {
	l, r := 0, 0
	response := []rune(s)
	for r < len(s) {
		if string(response[r]) != " " {
			r++
			continue
		}
		reverse(response[l:r])
		r++
		l = r
	}
	reverse(response[l:r])
	return string(response)
}

func reverse(sub []rune) {
	for i, j := 0, len(sub)-1; i < j; i, j = i+1, j-1 {
		sub[i], sub[j] = sub[j], sub[i]
	}
}

func main() {
	testcases := []string{"Let's take LeetCode contest", "Mr Ding"}
	for _, testcase := range testcases {
		fmt.Println(reverseWords(testcase))
	}
}
