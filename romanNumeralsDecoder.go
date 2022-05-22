package kata

func Decode(roman string) int {

	prev := byte('A')
	sum := 0

	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	for i := len(roman) - 1; i >= 0; i-- {
		if prev == 'A' {
			sum = sum + m[roman[i]]
		} else if m[prev] > m[roman[i]] {
			sum = sum - m[roman[i]]
		} else {
			sum = sum + m[roman[i]]
		}

		prev = roman[i]

	}

	return sum
}
