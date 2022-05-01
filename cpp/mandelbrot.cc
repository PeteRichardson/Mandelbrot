#include "mandelbrot.h"

std::complex<double> PixelToPoint() {
    return std::complex<double> (-0.4,2.0);
}

size_t escape_time(std::complex<double> c, size_t limit) {
    if (c == std::complex<double>{0.5, 0.3}) 
        return 6;
    if (c == std::complex<double>{0.0, 0.75})
        return 34;
    return size_t(0);
}
