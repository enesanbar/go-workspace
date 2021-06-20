package grains

import (
	"errors"
)

// Square returns the number of grains in a given chessboard square.
func Square(n int) (uint64, error) {

	if n <= 0 || n > 64 {
		return 0, errors.New("invalid input")
	}

	return 1 << (n - 1), nil
}

// Total returns the total number of grains on the chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
