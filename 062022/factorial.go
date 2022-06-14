package kata

func Factorial(n int) int {
	prod := 1
	for i := n; i > 1; i-- {
		prod = i * prod
	}
	return prod
}
