package turtle

import (
	"image/color"
	"math"
)

type Turtle struct {
	x, y       float32
	angle      float64
	draw       bool
	color      color.Color
	lines      [][4]float32
	step       int
	maxSteps   int
	stepDrawer StepDrawer
}

type Mover interface {
	Forward(distance float32)
	Left(angle float64)
	Right(angle float64)
	Up()
	Down()

	SetColor(c color.Color)
	IncStep()

	Step() int
	MaxSteps() int
}

type Drawer interface {
	NextStep()

	Lines() [][4]float32
	Color() color.Color
}

type StepDrawer interface {
	Step(turtle Mover)
}

var _ Mover = (*Turtle)(nil)

func NewTurtle(x, y float32, stepDrawer StepDrawer) *Turtle {
	return &Turtle{
		x:          x,
		y:          y,
		angle:      0,
		draw:       true,
		color:      color.White,
		step:       0,
		maxSteps:   360,
		stepDrawer: stepDrawer,
	}
}

func NewTurtleWithMaxSteps(x, y float32, maxSteps int, stepDrawer StepDrawer) *Turtle {
	return &Turtle{
		x:          x,
		y:          y,
		angle:      0,
		draw:       true,
		color:      color.White,
		step:       0,
		maxSteps:   maxSteps,
		stepDrawer: stepDrawer,
	}
}

func (t *Turtle) Color() color.Color {
	return t.color
}

func (t *Turtle) Lines() [][4]float32 {
	return t.lines
}

func (t *Turtle) IncStep() {
	t.step++
}

func (t *Turtle) SetColor(c color.Color) {
	t.color = c
}

func (t *Turtle) MaxSteps() int {
	return t.maxSteps
}

func (t *Turtle) Step() int {
	return t.step
}

func (t *Turtle) Up() {
	t.draw = false
}

func (t *Turtle) Down() {
	t.draw = true
}

func (t *Turtle) Forward(distance float32) {
	newX := t.x + float32(math.Cos(t.angle))*distance
	newY := t.y + float32(math.Sin(t.angle))*distance

	if t.draw {
		t.lines = append(t.lines, [4]float32{t.x, t.y, newX, newY})
	}

	t.x, t.y = newX, newY
}

func (t *Turtle) Left(angle float64) {
	t.angle += angle
}

func (t *Turtle) Right(angle float64) {
	t.angle -= angle
}

func (t *Turtle) NextStep() {
	if t.Step() > t.MaxSteps() {
		return
	}

	t.stepDrawer.Step(t)
}
