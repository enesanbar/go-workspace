package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(x int) bool
type unaryFunc func(x int) int

func (l IntList) Foldr(fn binFunc, initial int) int {
	for i := len(l) - 1; i >= 0; i-- {
		initial = fn(l[i], initial)
	}

	return initial
}

func (l IntList) Foldl(fn binFunc, initial int) int {
	for _, item := range l {
		initial = fn(initial, item)
	}

	return initial
}

func (l IntList) Filter(fn predFunc) IntList {
	result := IntList{}

	for _, item := range l {
		if fn(item) {
			result = append(result, item)
		}
	}

	return result
}

func (l IntList) Length() (length int) {
	for range l {
		length++
	}

	return length
}

func (l IntList) Map(fn unaryFunc) IntList {
	result := make([]int, l.Length())

	for i, item := range l {
		result[i] = fn(item)
	}

	return result
}

func (l IntList) Reverse() IntList {
	out := IntList{}

	for i := len(l) - 1; i >= 0; i-- {
		out = append(out, l[i])
	}

	return out
}

func (l IntList) Append(l2 IntList) IntList {
	res := make([]int, l.Length()+l2.Length())
	copy(res[:l.Length()], l)
	copy(res[l.Length():], l2)
	return res
}

func (l IntList) Concat(lists []IntList) IntList {
	out := l

	for _, list := range lists {
		out = out.Append(list)
	}

	return out
}
