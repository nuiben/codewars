package kata

func SpinWords(str string) string {

	var returnString string
	var revWord string
	var normalWord string

	if len(str) < 5 {
		return str
	}

	for _, i := range str {
		if i == ' ' && len(revWord) > 4 {
			returnString = returnString + revWord + " "
			revWord = ""
			normalWord = ""
		} else if i == ' ' && len(revWord) < 5 {
			returnString = returnString + normalWord + " "
			revWord = ""
			normalWord = ""
		} else {
			revWord = string(i) + revWord
			normalWord = normalWord + string(i)
		}
	}

	if len(revWord) > 4 {
		return returnString + revWord
	} else {
		return returnString + normalWord
	}
}
