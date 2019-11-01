package maths

func Factorial(n uint64) (fact uint64) {
	if n > 0 {
		fact = n * Factorial(n-1)
		return fact
	}
	return 1
}
