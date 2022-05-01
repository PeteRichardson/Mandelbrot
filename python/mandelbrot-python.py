import png

def pixel_to_point(bounds, pixel, upper_left, lower_right): 
    return 1 + 0.5j

def render(image, height, width, upper_left, lower_right):
    pass

def write_image(filename, image):
    png.from_array(image, 'L').save(filename)

if __name__ == "__main__":
    image = [[255, 0, 0, 255],
             [0, 255, 255, 0]]
    filename = "mandelbrot-python.png"
    height = 750
    width  = 1000
    upper_left = (-1, 1)
    lower_right = (1, -1)

    render(image, height, width, upper_left, lower_right)

    write_image(filename, image)
    