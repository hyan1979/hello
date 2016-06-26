package grammar

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() int
}

type Reactangle struct {
	width, height int
}

type Triangle struct {
	width, height int
}

type Circle struct {
	radius float64
}

func (r Reactangle) Area() int {
	return r.width * r.height
}

func (t Triangle) Area() int {
	return (t.width * t.height) / 2
}

func (c Circle) Area() float64 {
	return (c.radius * c.radius) * math.Pi
}

func GetArea() {
	r := Reactangle{3, 5}
	t := Triangle{3, 5}
	c := Circle{10.0}

	fmt.Println("Area of the Rectangle : ", r.Area())
	fmt.Println("Area of the Triangle : ", t.Area())
	fmt.Println("Area of the Circle : ", c.Area())

	s := Shaper(t)
	fmt.Println("Area of the Shaper(t) : ", s.Area())
}
