package pipe

import (
	"io"
	"strconv"
)

type Counter struct {
	Writer io.Writer
}

func (f *Counter) Count(n uint64) uint64 {
	if n == 0 {
		f.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}
	cur := n
	f.Writer.Write([]byte(strconv.FormatUint(cur, 10)))
	return f.Count(n - 1)
}
