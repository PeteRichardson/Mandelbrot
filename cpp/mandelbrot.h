#ifndef MANDELBROT_H
#define MANDELBROT_H

#include <complex>      // std::complex

std::complex<double> PixelToPoint();
size_t escape_time(std::complex<double>, size_t);
#endif