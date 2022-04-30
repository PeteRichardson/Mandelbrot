package main

import (
	"image"
	"image/png"
	"os"
)

type Point struct {
	x, y int
}

func escape_time(c complex128, limit uint8) uint8 {
	z := complex(0, 0)
	for i := uint8(0); i < limit; i++ {
		if (real(z)*real(z) + imag(z)*imag(z)) > 4.0 {
			return i
		}
		z = z*z + c
	}
	return 255
}

func pixel_to_point(width int, height int, row int, column int, upper_left complex128, lower_right complex128) complex128 {
	var pixwidth float64 = real(lower_right) - real(upper_left)
	var pixheight = imag(upper_left) - imag(lower_right)
	var resreal = real(upper_left) + (float64(row))*pixwidth/float64(width)
	var resimag = imag(upper_left) - (float64(column))*pixheight/float64(height)
	return complex(resreal, resimag)
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
	const WIDTH int = 4000
	const HEIGHT int = 3000

	myImage := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	render(myImage, WIDTH, HEIGHT, -1.20+0.35i, -1+0.2i)

	write_image("test.png", myImage)
}
