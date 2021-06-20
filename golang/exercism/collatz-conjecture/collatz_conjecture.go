package collatzconjecture

import "github.com/pkg/errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Negative numbers not allowed")
	}

	steps := 0

	for n != 1 {
		n = [2]int{n >> 1, 3*n + 1}[n&1]
		steps++
	}

	return steps, nil
}
