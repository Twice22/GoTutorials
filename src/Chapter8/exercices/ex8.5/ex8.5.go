package main 

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"runtime"
	"time"
	"fmt"
	"sync"
	"net/http"
	"log"
)

// go run ex8.5.go
func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height		   = 1024, 1024
	)

	// number of worker = number of CPU on the current machine
	var workers int = runtime.GOMAXPROCS(runtime.NumCPU())

	// Timer
	start := time.Now()

	// number of working goroutines
	var wg sync.WaitGroup

	// create unbuffered channel
	rows := make(chan int, height)

	// send 0, 1, 2,... up to height of the  rows channel
	for row := 0; row < height; row++ {
		rows <- row
	}
	close(rows) // close rows channel

	// create new 1024x1024 rectangle
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// for every CPU
	for i := 0; i < workers; i++ {
		wg.Add(1) // add one goroutine to the active goroutines

		go func() {

			defer wg.Done() // remove one goroutine from active goroutines when fct end

			// for each row in rows channel that are not yet processed
			for py := range rows {
				// retrieve y coord in [-2,2] from py in [0, 1024[
				y := float64(py) / height*(ymax-ymin) + ymin

				for px := 0; px < width; px++ {

					// retrieve x coord in [-2,2] from px in [0, 1024[
					x := float64(px) / width * (xmax-xmin) + xmin

					// z = x + iy
					z := complex(x, y)

					// pixel (px, py) represents complex value z
					img.Set(px, py, mandelbrot(z))
				}
			}
		}()
	}

	// wait until all row from rows channel are been processed (received and processed through range loop)
	wg.Wait()

	fmt.Println("elapsed time: ", time.Since(start))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img) // NOTE: ignoring errors
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	// test whether repeatedly squarring and adding the number "escapes" the circle
	// of radius 2. If so, the point is shaded (see color.Gray) by the number of iterations
	// it took to escape
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}