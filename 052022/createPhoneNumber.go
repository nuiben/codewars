package kata

import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {

	var phoneNum string = "("

	for k, i := range numbers {
		if k == 3 {
			phoneNum = phoneNum + ") "
		} else if k == 6 {
			phoneNum = phoneNum + "-"
		}
		phoneNum = phoneNum + fmt.Sprint(i)
	}

	return phoneNum
}
