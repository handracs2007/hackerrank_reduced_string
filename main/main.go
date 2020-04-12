// https://www.hackerrank.com/challenges/reduced-string/problem

package main

import "fmt"

func superReducedString(s string) string {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			var temp = ""
			if i-1 > 0 {
				temp += s[0 : i-1]
			}

			if i+1 < len(s) {
				temp += s[i+1:]
			}

			s = temp

			if i == 1 {
				i = 0
			} else {
				i -= 2
			}
		}
	}

	if len(s) == 0 {
		return "Empty String"
	}

	return s
}

func main() {
	fmt.Println(superReducedString("aaabccddd "))
	fmt.Println(superReducedString("baab"))
}
