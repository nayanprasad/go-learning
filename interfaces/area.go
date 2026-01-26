package main

import (
	"fmt"
)

type shape interface {
	area() float64
}

type triangle struct {
	heigth float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 3, heigth: 5}

	fmt.Println(getArea(t))

	s := square{sideLength: 5}

	fmt.Println(getArea(s))

}

func getArea(s shape) float64 {
	return s.area()
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.heigth
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}
