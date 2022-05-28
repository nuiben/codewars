package kata

import "strings"

func ToJadenCase(str string) string {

	var word string
	var jadenStr string
	var counter int = 0

	for _, i := range str {
		if i == ' ' {
			jadenStr = jadenStr + word + " "
			word = ""
			counter = 0
		} else if counter == 0 {
			word = word + strings.ToUpper(string(i))
			counter++
		} else {
			word = word + string(i)
			counter++
		}
	}
	return jadenStr + word
}
