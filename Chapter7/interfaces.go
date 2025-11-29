package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	w, h float64
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.w + r.h)
}

type MultiShape struct {
	shapes []Shape
}

func total(shapes []Shape, f func(Shape) float64) float64 {
	var sum float64
	for _, s := range shapes {
		sum += f(s)
	}
	return sum
}

func main() {
	multiShape := MultiShape{
		shapes: []Shape{
			Circle{5},
			Rectangle{10, 10},
		},
	}

	fmt.Println("Total Area:", total(multiShape.shapes, Shape.area))
	fmt.Println("Total Perimeter:", total(multiShape.shapes, Shape.perimeter))
}
