package kata

import "fmt"

func Solution(word string) string {
	reverse := ""
	for _, i := range word {
		reverse = fmt.Sprintf("%c", i) + reverse
	}
	return reverse
}
