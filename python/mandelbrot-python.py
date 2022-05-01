from os import lseek
import png


def render(image):
    image = image = [[255, 0, 0, 255],
             [0, 255, 255, 0]]

def write_image(filename, image):
    png.from_array(image, 'L').save(filename)

if __name__ == "__main__":

    filename = "mandelbrot-python.png"
    height = 
    render(image, height, width, upper_left, lower_right)

    write_image(filename, image)
    
    print("Hello, Carl!\n")