package kata

func alphanumeric(str string) bool {

	if len(str) == 0 {
		return false
	}

	for _, i := range str {
		if !((i > 47 && i < 58) || (i > 64 && i < 91) || (i > 96 && i < 123)) {
			return false
		}
	}
	return true
}
