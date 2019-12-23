package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{height: 10, base: 10}
	squ := square{sideLength: 10}

	printArea(tri)
	printArea(squ)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
