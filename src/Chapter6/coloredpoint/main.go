package main

import (
	"fmt"
	"math"
)

import "image/color"

type Point struct{X, Y float64}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}


func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	// we can call method of the embedded Point field
	// using a receiver (p here) of type ColoredPoint.
	// Note: we still need to select the Point field of q
	// by using q.Point as a parameter to Point.Distance method!!
	// p.Distance(q) wouldn't work!
	fmt.Println(p.Distance(q.Point)) // "5"

	p.ScaleBy(2)
	q.ScaleBy(2)

	fmt.Println(p.Distance(q.Point)) // "10"
}

// RECALL: init function are called first when running the program
func init() {
	p := Point{1, 2}
	q := Point{4, 6}

	// the selector p.Distance yields a method value, a function
	// that binds a method (Point.Distance) to a specific receiver
	// value p.
	distanceFromP := p.Distance // method value
	fmt.Println(distanceFromP(q)) // "5"
	var origin Point
	fmt.Println(distanceFromP(origin)) // sqrt(5)

	scaleP := p.ScaleBy
	scaleP(2) // p becomes (2, 4)
	scaleP(3) // p becomes (6, 12)
	scaleP(10) // p becomes (60, 120)
}

func init() {
	p := Point{1, 2}
	q := Point{4, 6}

	// a method expression, written T.f or (*T).f where T
	// is a type, yields a function value with a regular first parameter
	// taking the place of the receiver, so it can be call in the traditional way.
	distance := Point.Distance // method expression
	fmt.Println(distance(p, q)) // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p) // "{2, 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"

}

func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	// the type to an anymous field can be a pointer to a named type
	// (here *Point).
	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point // p and q now share the same Point
	p.ScaleBy(2)
	// NOTE: *p.Point means *(p.Point). (*p).Point
	// as no sense here since p is a named type (not a pointer type)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
}