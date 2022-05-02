#ifndef MANDELBROT_H
#define MANDELBROT_H

#include <complex>      // std::complex

std::complex<double> PixelToPoint(int, int, int, int, std::complex<double>, std::complex<double>);
size_t escape_time(std::complex<double>, size_t);
#endif