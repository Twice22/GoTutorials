package main 

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320		// canvas size in pixels
	cells = 100						// number of grid cells
	xyrange = 30.0					// axis ranges (-xyrange..+xyrange)
	xyscale = width / 2 / xyrange	// pixels per x or y unit
	zscale = height * 0.4			// pixel per z unit
	angle = math.Pi / 6				// angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	// create a svg canvas of size 600x320
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, h := corner(i+1, j+1)

			// I wanted to create a gradient but It would have been that I need to retrieve the max
			// and the min in a loop and then compute the percentage of RGB accordingly to the max
			// of each layers. As I didn't do that, the gradient is far from being perfect.
			var color = fmt.Sprintf("style='fill:rgb(%d%%, 0%%, %d%%)'", (50+h)%100, (50-h)%100)

			// draw a polygon that links all the points ax to dy
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' %s/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, int) {
	// retrieve (x,y) coords of corner (i,j) 
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)

	// Compute surface height z
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y) * cos30 * xyscale
	sy := height/2 + (x+y) * sin30 * xyscale - z*zscale

	return sx, sy, int(100 * z)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0) (Hypot returns Sqrt(x*x + y*y))
	if r != 0 {
		return math.Sin(r) / r	
	}
	return 0
}