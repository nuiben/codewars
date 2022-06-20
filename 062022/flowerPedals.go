package kata

func HowMuchILoveYou(i int) string {
	m := make(map[int]string)
	m[1] = "I love you"
	m[2] = "a little"
	m[3] = "a lot"
	m[4] = "passionately"
	m[5] = "madly"
	m[0] = "not at all"
	return m[i%6]
}
