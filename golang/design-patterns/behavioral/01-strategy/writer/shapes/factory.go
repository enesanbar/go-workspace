package shapes

import (
	"fmt"
	"os"

	"github.com/enesanbar/workspace/golang/design-patterns/behavioral/01-strategy/writer"
)

const (
	TextStrategy  = "text"
	ImageStrategy = "image"
)

func Factory(s string) (writer.Output, error) {
	switch s {
	case TextStrategy:
		return &TextSquare{
			DrawOutput: writer.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case ImageStrategy:
		return &ImageSquare{
			DrawOutput: writer.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}
