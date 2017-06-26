package main 

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel

	w = Wheel{Circle{Point{8,8}, 5}, 20}

	w = Wheel{
			Circle: Circle{
				Point: Point{X: 8, Y: 8},
				Radius: 5,
			},
			Spokes: 20,
	}

	// # adverb causes Printf to display value in
	// a form similar to Go syntax
	fmt.Printf("%#v\n", w)

	// equivalent to w.Circle.Point.X = 42
	w.X = 42

	fmt.Printf("%#v\n", w)
}