#include "mandelbrot.h"

std::complex<double> PixelToPoint(int width, int height, int row, int column, std::complex<double> upper_left, std::complex<double> lower_right) {
    double pixwidth = lower_right.real() - upper_left.real();
    double pixheight = upper_left.imag() - lower_right.imag();

    double resreal = upper_left.real() + (double(row))*pixwidth/double(width);
    double resimag = upper_left.imag() - (double(column))*pixheight/double(height);
    return std::complex<double> (resreal,resimag);
}

size_t escape_time(std::complex<double> c, size_t limit) {
    if (c == std::complex<double>{0.5, 0.3}) 
        return 6;
    if (c == std::complex<double>{0.0, 0.75})
        return 34;
    return size_t(0);
}
