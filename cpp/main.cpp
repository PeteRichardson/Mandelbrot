#include <iostream>
#include <string>
#include <vector>
#include <complex>
#include "lodepng.h"
#include "mandelbrot.h"

using Image = std::vector<unsigned char>;

void render(Image& image, unsigned width, unsigned height) {
    for (int row=0; row < width; row++) {
        for (int column=0; column < height; column++) {
            image[row * width + column] = 128;
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
    Image image(width*height*4);
    render(image, width, height);
    write_file("mandelbrot-cpp.png", image, width, height);
    return 0;
}


