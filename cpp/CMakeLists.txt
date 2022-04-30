CMAKE_MINIMUM_REQUIRED(VERSION 3.18)

project ( mandelbrot-cpp)

add_executable(mandelbrot-cpp main.cpp lodepng.cpp)

target_compile_features(${PROJECT_NAME} PUBLIC cxx_std_20)
set (CMAKE_EXPORT_COMPILE_COMMANDS true)
