package samples

import (
	"github.com/kodlan/gograph/internal/tcolor"
	"github.com/kodlan/gograph/internal/turtle"
	"math"
)

type Stars struct {
	starsDrawn int
	maxStars   int
	edge       float32
}

func (s *Stars) Step(turtle turtle.Mover) {

	if s.starsDrawn >= s.maxStars {
		return // nothing more to do
	}

	// Draw one edge of the 5‑point star.
	turtle.Forward(s.edge)
	turtle.Right(144 * math.Pi / 180) // 144° makes a regular star

	// After 5 edges the star is complete.
	if turtle.Step()%6 == 0 {
		s.starsDrawn++

		// Shrink the next star and shift colour.
		s.edge *= 0.88
		hue := float64(s.starsDrawn) / float64(s.maxStars)
		turtle.SetColor(tcolor.Hsv(hue))
	}
}

func NewStars() *Stars {
	return &Stars{
		starsDrawn: 0,
		maxStars:   30,
		edge:       300,
	}
}
