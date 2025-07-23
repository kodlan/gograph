package turtle

import (
	"image/color"
	"math"
)

type Turtle struct {
	x, y       float32
	angle      float32
	draw       bool
	color      color.Color
	lines      [][4]float32
	step       int
	maxSteps   int
	stepDrawer StepDrawer
}

/*
Mover is an interface that provides methods to draw with the turtle, like move forward,
turn left, right and so on. Used mostly in concrete implementations.
*/
type Mover interface {
	Forward(distance float32)

	Left(rad float32)
	LeftAngle(angle float32)
	Right(rad float32)
	RightAngle(angle float32)

	Up()
	Down()

	SetColor(c color.Color)
	IncStep()

	Step() int
	MaxSteps() int
}

/*
Drawer interface is used to actually render the drawing on the screen with some specific renderer.
Provides access to lines using Lines getter and to the method that generates next step - NextStep
*/
type Drawer interface {
	NextStep()

	Lines() [][4]float32
	Color() color.Color
}

/*
StepDrawer provides only one method to generate the next step of the turtle drawing.
Should be implemented in concrete realisation.
*/
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
	newX := t.x + float32(math.Cos(float64(t.angle)))*distance
	newY := t.y + float32(math.Sin(float64(t.angle)))*distance

	if t.draw {
		t.lines = append(t.lines, [4]float32{t.x, t.y, newX, newY})
	}

	t.x, t.y = newX, newY
}

func (t *Turtle) Left(angle float32) {
	t.angle += angle
}

func (t *Turtle) Right(angle float32) {
	t.angle -= angle
}

func (t *Turtle) NextStep() {
	if t.Step() > t.MaxSteps() {
		return
	}

	t.stepDrawer.Step(t)
	t.IncStep()
}

func AngleToRad(angle float32) float32 {
	return angle * math.Pi / 180.0
}

func (t *Turtle) LeftAngle(angle float32) {
	t.Left(AngleToRad(angle))
}

func (t *Turtle) RightAngle(angle float32) {
	t.Right(AngleToRad(angle))
}
