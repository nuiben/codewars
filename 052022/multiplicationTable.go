package kata

import "fmt"

func MultiTable(number int) string {
	var table = ""
	for i := 1; i <= 10; i++ {
		if i != 10 {
			table = table + fmt.Sprint(i) + " * " + fmt.Sprint(number) + " = " + fmt.Sprint(i*number) + "\n"
		} else {
			table = table + fmt.Sprint(i) + " * " + fmt.Sprint(number) + " = " + fmt.Sprint(i*number)
		}

	}
	return table
}
