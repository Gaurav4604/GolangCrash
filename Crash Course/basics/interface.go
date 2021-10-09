package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type shape_2 interface {
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// both rect and circle have to implement methods of shape
// if they do, then they belong to the type shape

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

func (r *rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 5}

	shapes := []shape{&r, c}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

	shapes_2 := []shape_2{&r, c}
	for _, shape_2 := range shapes_2 {
		fmt.Println(shape_2.perimeter())
	}
}
