package kata

import (
	"strings"
)

func IsPalindrome(str string) bool {
	if len(str) == 1 {
		return true
	}
	for h, t := 0, len(str)-1; h < t; h, t = h+1, t-1 {
		if !strings.EqualFold(string(str[h]), string(str[t])) {
			return false
		}
	}

	return true
}
