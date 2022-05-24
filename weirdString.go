package kata

import (
	"strings"
)

func toWeirdCase(str string) string {

	var word string
	var weirdStr string
	var counter int = 0

	for _, i := range str {
		if i == ' ' {
			weirdStr = weirdStr + word + " "
			word = ""
			counter = 0
		} else if counter%2 == 0 {
			word = word + strings.ToUpper(string(i))
			counter++
		} else {
			word = word + strings.ToLower(string(i))
			counter++
		}

	}

	return weirdStr + word

}
