#include <iostream>
#include <iomanip>
#include <string>
#include <vector>
#include <complex>
#include "lodepng.h"
#include "mandelbrot.h"

using Image = std::vector<unsigned char>;

void render(Image& image, unsigned width, unsigned height, std::complex<double> upper_left, std::complex<double> lower_right) {
    for (int row=0; row < width; row++) {
        for (int column=0; column < height; column++) {
            int i = row*width + column;
			std::complex<double> point = PixelToPoint(width, height, row, column, upper_left, lower_right);
			uint8_t color = 255 - escape_time(point, 255);
			image[i*4+0] = color;
			image[i*4+1] = color;
			image[i*4+2] = color;
			image[i*4+3] = 255;
        }
    }
}

void write_file(const char* filename, Image& image, unsigned width, unsigned height) {
  unsigned error = lodepng::encode(filename, image, width, height);
  if(error) std::cerr << "# encoder error " << error << ": "<< lodepng_error_text(error) << std::endl;
}

int main(int argc, const char * argv[]) {
    unsigned width =  1000;
    unsigned height = 750;
    std::complex<double> upper_left(-1.20, 0.35);
    std::complex<double> lower_right(-1.0, 0.20);
    Image image(width*height*4);
    std::cout << "# Lang=Cpp,  Width=" << std::setw(5) << width << ", Height=" << std::setw(5) << height << ", upper_left=" << upper_left << ", lower_right=" << lower_right << std::endl;

    render(image, width, height, upper_left, lower_right);
    write_file("mandelbrot-cpp.png", image, width, height);
    return 0;
}


