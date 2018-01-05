package isPalindrome

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	sarr := make([]string, 0)
	for _, r := range s {
		sarr = append(sarr, string(r))
	}
	fmt.Println(len(sarr))
	for i := range sarr {
		if sarr[i] != sarr[len(sarr)-1-i] {
			return false
		}
	}

	return true
}
