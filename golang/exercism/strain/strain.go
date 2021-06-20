// Package strain implements Keep and Discard operations on collections.
package strain

type Ints []int
type Strings []string
type Lists [][]int

func (numbers Ints) Keep(fn func(int) bool) Ints {
	var result Ints

	for _, number := range numbers {
		if fn(number) {
			result = append(result, number)
		}
	}

	return result
}

func (numbers Ints) Discard(fn func(int) bool) Ints {
	return numbers.Keep(func(n int) bool {
		return !fn(n)
	})
}

func (strings Strings) Keep(fn func(string) bool) Strings {
	var result Strings

	for _, s := range strings {
		if fn(s) {
			result = append(result, s)
		}
	}

	return result
}

func (lists Lists) Keep(fn func([]int) bool) Lists {
	var result Lists

	for _, list := range lists {
		if fn(list) {
			result = append(result, list)
		}
	}

	return result
}
