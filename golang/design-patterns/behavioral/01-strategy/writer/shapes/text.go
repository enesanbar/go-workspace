package shapes

import "github.com/enesanbar/workspace/golang/design-patterns/behavioral/01-strategy/writer"

type TextSquare struct {
	writer.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
