// This problem considers that a valid parentheses has an open (+1) and close (-1) which will net to
// a value of 0 if valid. Because close parentheses can be input before an open parentheses,
// we return false if the net value ever becomes negative.

package kata

func ValidParentheses(parens string) bool {

	net := 0

	for _, i := range parens {
		if i == 40 { //40 = open paren
			net++
		} else if i == 41 { //41 = close paren
			net--
			if net < 0 {
				return false
			}
		}
	}

	return net == 0
}
