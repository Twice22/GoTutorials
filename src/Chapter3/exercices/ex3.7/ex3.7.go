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
		y1 := float64(py) / height*(ymax-ymin) + ymin

		// retrieve y+1
		y2 := float64(py+1) / height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {

			// retrieve x coord in [-2,2] from px in [0, 1024[
			x1 := float64(px) / width * (xmax-xmin) + xmin

			// retrieve y+1
			x2 := float64(px+1) / width * (xmax-xmin) + xmin

			// 4 different values for z (x1,y1), (x1, y2), (x2, y1), (x2, y2)
			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)

			// subsampling (taking average of the 4 points)
			// pixel (px, py) represents complex value z
			img.Set(px, py, newtonfrac((z1 + z2 + z3 + z4) / 4))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func newtonfrac(v complex128) color.Color {
	const iterations = 68
	const contrast = 3

	// fractal for z^4 - 1 = 0 equation
	// we use Newton's iteration method:
	// z^4 - 1 = 0
	// using Newton's iteration method:
	// z_{n+1} = z_n - p(z_n)/p'(z_n)
	// z_{n+1} = z_n - (z_n^4 - 1)/(4z_n^3)
	// so finally (using v instead of z_n)
	// v = v - (v^4 - 1)/(4v^3) = (3v^4 + 1)/(4v^3)

	for n := uint8(0); n < iterations; n++ {
		v = (3 * v*v*v*v + 1)/(4*v*v*v)
		if cmplx.Abs(v*v*v*v - 1) < 1e-4 {
			if real(v) > 0 && imag(v) > 0 {
				return color.RGBA{127 - contrast*n, 0, 0, 255}
			} else if real(v) > 0 && imag(v) < 0 {
				return color.RGBA{0, 127 - contrast*n, 0, 255}
			} else if real(v) < 0 && imag(v) > 0 {
				return color.RGBA{0, 0, 127 - contrast*n, 255}
			} else {
				return color.RGBA{127 - contrast*n, 127 - contrast*n, 127 - contrast*n, 255}
			}
		}
	}
	return color.RGBA{255,0,0,255}
}