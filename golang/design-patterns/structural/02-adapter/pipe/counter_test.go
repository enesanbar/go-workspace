package pipe

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCounter_Count(t *testing.T) {
	counter := Counter{
		Writer: os.Stdout,
	}
	counter.Count(5)
}

func TestCounter_Pipe(t *testing.T) {
	pipeReader, pipeWriter := io.Pipe()
	defer pipeReader.Close()
	defer pipeWriter.Close()

	counter := Counter{
		Writer: pipeWriter,
	}

	var bufferRead bytes.Buffer
	tee := io.TeeReader(pipeReader, &bufferRead)

	go func() {
		io.Copy(os.Stdout, tee)
	}()

	counter.Count(10)
	fmt.Println(bufferRead.String())
}
