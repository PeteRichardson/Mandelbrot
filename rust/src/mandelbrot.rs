use num::Complex;

pub fn escape_time(c: Complex<f64>, limit: usize) -> Option<usize> {
    let mut z = Complex { re: 0.0, im: 0.0 };
    for i in 0..limit {
        if z.norm_sqr() > 4.0 {
            return Some(i);
        }
        z = z * z + c
    }
    None
}

#[test]
fn test_escape_time() {
    assert_eq!(escape_time(Complex { re: 0.0, im: 0.0 }, 255), None);
    assert_eq!(escape_time(Complex { re: 0.5, im: 0.3 }, 255), Some(6));
    assert_eq!(escape_time(Complex { re: -0.5, im: 0.25 }, 255), None);
    assert_eq!(escape_time(Complex { re: 0.0, im: 0.75 }, 255), Some(34));
}

pub fn pixel_to_point( bounds : (usize, usize),
                    pixel: (usize, usize),
                    upper_left: Complex<f64>,
                    lower_right: Complex<f64>) 
                -> Complex<f64>
{
    let (width, height) = (lower_right.re - upper_left.re,
                            upper_left.im - lower_right.im);
    Complex {
        re: upper_left.re + pixel.0 as f64 * width /bounds.0 as f64,
        im: upper_left.im - pixel.1 as f64 * height/bounds.1 as f64
    }
}

#[test]
fn test_pixel_to_point() {
    assert_eq!(pixel_to_point((100,200), (25,175),
                                Complex { re: -1.0, im:  1.0},
                                Complex { re:  1.0, im: -1.0}),
                Complex{ re: -0.5, im: -0.75});
}
