package isPalindrome

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for i := range s {
		if string(s[i]) != string(s[len(s)-1-i]) {
			return false
		}
	}
	return true
}
