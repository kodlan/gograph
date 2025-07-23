package samples

import (
	"github.com/kodlan/gograph/internal/tcolor"
	"github.com/kodlan/gograph/internal/turtle"
)

type SquareSpiral struct {
	n float32
}

func (s *SquareSpiral) Step(turtle turtle.Mover) {
	h := float64(turtle.Step()) / float64(turtle.MaxSteps())
	turtle.SetColor(tcolor.Hsv(h))

	turtle.Forward(s.n)
	turtle.LeftAngle(89.5)

	s.n += 2
}

func NewSquareSpiral() *SquareSpiral {
	return &SquareSpiral{}
}
