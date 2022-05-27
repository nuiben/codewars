package kata

import "strconv"

func Is_valid_ip(ip string) bool {

	var perCounter int = 0 //period counter
	var octet string

	//by default, any entry that does not have 3 periods and 4 digits is excluded
	if len(ip) < 7 {
		return false
	}

	for k, i := range ip {
		if i == '.' {
			perCounter++
			if !(evaluate(octet)) {
				return false
			} // if the evaulate func failed, ip address no good
			octet = ""
		} else if k == len(ip)-1 { //when we reach the last octet
			if perCounter != 3 {
				return false
			}
			octet = octet + string(i)
			if !(evaluate(octet)) {
				return false
			}
		} else {
			octet = octet + string(i)
		}
	}
	return true
}

func evaluate(octet string) bool {

	//this handles cases where octet has a leading '0'
	if len(octet) > 1 && string(octet[0]) == "0" {
		return false
	}

	//convert ASCII to int (A to i), if the format is not a number, return false
	var octInt, err = strconv.Atoi(octet)
	if err != nil {
		return false
	}

	//valid number check
	if octInt >= 0 && octInt < 256 {
		return true
	}

	//default if the octet didn't pass
	return false
}
