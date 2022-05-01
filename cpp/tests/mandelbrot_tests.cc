#include <gtest/gtest.h>

#include "../mandelbrot.h"

TEST(MandelbrotTests, Escape_Time_Basic) {
    EXPECT_EQ(escape_time(std::complex<double> (0.0, 0.0 ), 255), 0);
    EXPECT_EQ(escape_time(std::complex<double> (0.5, 0.3 ), 255), 6);
    EXPECT_EQ(escape_time(std::complex<double> (-0.5,0.25), 255), 0);
    EXPECT_EQ(escape_time(std::complex<double> (0.0, 0.75), 255), 34);
}
