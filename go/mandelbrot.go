package main

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
