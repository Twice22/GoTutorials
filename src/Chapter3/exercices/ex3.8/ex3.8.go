package main 

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height		   = 1024, 1024
	)

	// create new 1024x1024 rectangle
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {


		// retrieve y coord in [-2,2] from py in [0, 1024[
		y := float64(py) / height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {

			// retrieve x coord in [-2,2] from px in [0, 1024[
			x := float64(px) / width * (xmax-xmin) + xmin

			// 4 different values for z (x1,y1), (x1, y2), (x2, y1), (x2, y2)
			z := complex(x, y)

			// subsampling (taking average of the 4 points)
			// pixel (px, py) represents complex value z
			img.Set(px, py, mandelbrotComplex64(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrotComplex64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64

	// test whether repeatedly squarring and adding the number "escapes" the circle
	// of radius 2. If so, the point is shaded (see color.Gray) by the number of iterations
	// it took to escape
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}