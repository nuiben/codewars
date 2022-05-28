package kata

type MyString string

func (s MyString) IsUpperCase() bool {
	// ASCII values > 90 are lowercase
	for _, i := range s {
		if i > 90 {
			return false
		}
	}
	return true
}
