package samples

import (
	"github.com/kodlan/gograph/internal/tcolor"
	"github.com/kodlan/gograph/internal/turtle"
	"math"
)

type Spiral struct {
}

func (s *Spiral) Step(turtle turtle.Mover) {
	h := float64(turtle.Step()) / float64(turtle.MaxSteps())
	turtle.SetColor(tcolor.Hsv(h))

	turtle.Forward(0.3 + float32(turtle.Step())*0.3)
	turtle.Left(math.Pi / 6.)
}

func NewSpiral() *Spiral {
	return &Spiral{}
}
