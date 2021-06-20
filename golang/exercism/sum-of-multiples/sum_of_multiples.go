package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if d > 0 && i%d == 0 {
				sum += i
				break
			}
		}
	}

	return
}
