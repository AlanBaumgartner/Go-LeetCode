package strStr

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	index := 0
	buffer := 0

	if needle == "" {
		return 0
	}
	if strings.Contains(haystack, needle) {
		for index+buffer < len(haystack) && buffer < len(needle) {
			if string(haystack[index+buffer]) != string(needle[buffer]) {
				buffer = 0
				index++
			} else {
				buffer++
			}
		}
		return index
	}
	return -1
}
