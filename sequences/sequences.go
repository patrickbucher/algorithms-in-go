package sequences

func Fibs(n int) []int {
	if n <= 0 {
		return []int{}
	}
	a, b := 1, 1
	fibs := make([]int, n)
	for i := 0; i < n; i++ {
		fibs[i] = a
		c := a + b
		a = b
		b = c
	}
	return fibs
}

func Primes(l int) []int {
	primes := make([]int, 0)
	for i := 2; i <= l; i++ {
		prime := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			primes = append(primes, i)
		}
	}
	return primes
}

func PrimeSieve(l int) []int {
	primes := make([]int, 0)
	for i := 2; i <= l; i++ {
		prime := true
		for _, p := range primes {
			if i%p == 0 {
				prime = false
				break
			}
		}
		if prime {
			primes = append(primes, i)
		}
	}
	return primes
}
