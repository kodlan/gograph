package turtle

import (
	"image/color"
	"math"
)

type Turtle struct {
	x, y  float64
	angle float64
	draw  bool
	color color.Color
	lines [][4]float64
}

type Mover interface {
	Forward(distance float64)
	Up()
	Down()
	Left(angle float64)
	Right(angle float64)
}

var _ Mover = (*Turtle)(nil)

func NewTurtle(x, y float64) *Turtle {
	return &Turtle{
		x:     x,
		y:     y,
		angle: 0,
		draw:  true,
		color: color.White,
	}
}

func (t *Turtle) Up() {
	t.draw = false
}

func (t *Turtle) Down() {
	t.draw = true
}

func (t *Turtle) Forward(distance float64) {
	newX := t.x + math.Cos(t.angle)*distance
	newY := t.y + math.Sin(t.angle)*distance

	if t.draw {
		t.lines = append(t.lines, [4]float64{t.x, t.y, newX, newY})
	}

	t.x, t.y = newX, newY
}

func (t *Turtle) Left(angle float64) {
	t.angle += angle
}

func (t *Turtle) Right(angle float64) {
	t.angle -= angle
}
