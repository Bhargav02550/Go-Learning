package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	var shape Shape

	shape = Rectangle{
		Width:  10,
		Height: 20,
	}

	fmt.Println(shape.Area())

	shape = Circle{
		Radius: 2,
	}

	fmt.Println(shape.Area())
}
