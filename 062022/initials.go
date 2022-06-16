package kata

import (
	"strings"
)

func AbbrevName(name string) string {
	initials := ""
	space := false
	for k, i := range name {
		if k == 0 {
			initials += strings.ToUpper(string(i)) + "."
		} else if i == 32 {
			space = true
		} else if i != 32 && space == true {
			initials += strings.ToUpper(string(i))
			space = false
		}
	}
	return initials
}
