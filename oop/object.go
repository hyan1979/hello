package oop

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Max() float64 {
	return math.Max(v.X, v.Y)
}

func Vertex_method() {
	v := &Vertex{3, 4}
	fmt.Println("Abs() = ", v.Abs())
	fmt.Println("Max() = ", v.Max())
}

type entity float32

func (e *entity) inc() {
	(*e)++
}

func (e entity) echo() {
	fmt.Println(e)
}

func Entity_method() {
	var e entity = 3
	e.inc()
	e.echo()
	e.inc()
	e.echo()
}
