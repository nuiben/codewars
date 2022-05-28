package kata

import (
	"strings"
)

func GetCount(str string) (count int) {
	s := "aeiou"
	sum := 0
	for i := 0; i < len(str); i++ {
		tp := string(str[i])
		if strings.Contains(s, tp) {
			sum++
		}
	}
	return sum
}
