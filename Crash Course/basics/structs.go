package main

import "fmt"

// type tells that the given type is a custom type
// struct tells that it is a structure
type Point struct {
	x int
	y int
}

type Circle struct {
	radius float64
	center *Point
}

func changeX(p *Point) {
	(*p).x = 100
}

// binding the method to the struct
func (c Circle) getRadius() float64 {
	return c.radius
}

func (c *Circle) setRadius(newRadius float64) {
	c.radius = newRadius
}

func main() {
	p1 := Point{10, 12}
	p2 := Point{12, 14}

	p3 := Point{x: 1}

	fmt.Println(p1)

	changeX(&p1)

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	circle := Circle{radius: 5.6, center: &Point{x: 10, y: 20}}
	fmt.Println(circle)
	fmt.Println(circle.center.x)
	fmt.Println(circle.getRadius())

	circle.setRadius(10.2)
	fmt.Println(circle.getRadius())
}
