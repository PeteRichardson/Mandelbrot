package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

type Point struct {
	x, y int
}

func render(myImage *image.RGBA, width int, height int, upper_left complex128, lower_right complex128) {
	if len(myImage.Pix)/4 != width*height {
		panic("render failure... Pix size doesn't match bounds")
	}

	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			i := row*width + column
			var point complex128 = pixel_to_point(width, height, row, column, upper_left, lower_right)
			var color uint8 = 255 - escape_time(point, 255)
			myImage.Pix[i*4+0] = color
			myImage.Pix[i*4+1] = color
			myImage.Pix[i*4+2] = color
			myImage.Pix[i*4+3] = 255 // alpha
		}
	}
}

func write_image(filename string, myImage image.Image) {
	outputFile, err := os.Create(filename)
	if err != nil {
		// Handle error
	}
	png.Encode(outputFile, myImage)
	outputFile.Close()
}

func main() {
	const WIDTH int = 1000
	const HEIGHT int = 750

	myImage := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	const upper_left = -1.20 + 0.35i
	const lower_right = -1 + 0.2i

	fmt.Printf("# Lang=Go,   Width=%5d, Height=%5d, upper_left=%v, lower_right=%v\n", WIDTH, HEIGHT, upper_left, lower_right)
	render(myImage, WIDTH, HEIGHT, upper_left, lower_right)

	write_image("mandelbrot-go.png", myImage)
}
