package main 

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
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
			img.Set(px, py, mandelbrotBigFloat(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	zr := (&big.Float{}).SetFloat64(real(z))
	zi := (&big.Float{}).SetFloat64(imag(z))

	vr, vi := &big.Float{}, &big.Float{}

	for n := uint8(0); n < iterations; n++ {
		// real part of v*v (vrvr): vr*vr - vi*vi
		vrvr := &big.Float{}
		vrvr.Mul(vr, vr).Sub(vrvr, (&big.Float{}).Mul(vi, vi))

		// imag part of v*v (vivi): 2 * vr * vi
		vivi := &big.Float{}
		vivi.Mul(vi, vr).Mul(vivi, big.NewFloat(2))

		vr.Add(vrvr, zr)
		vi.Add(vivi, zi)

		testr := (&big.Float{}).Mul(vr, vr)
		testi := (&big.Float{}).Mul(vi, vi)

		test := (&big.Float{})
		test.Add(testr, testi)

		// return 1 if test > 4
		if test.Cmp(big.NewFloat(4)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}