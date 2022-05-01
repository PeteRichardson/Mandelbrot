CMAKE_MINIMUM_REQUIRED(VERSION 3.18)

include(FetchContent)
FetchContent_Declare(
    googletest
    GIT_REPOSITORY https://github.com/google/googletest
    GIT_TAG        release-1.11.0
)
FetchContent_MakeAvailable(googletest)

project ( mandelbrot-cpp)

add_executable(mandelbrot-cpp main.cpp lodepng.cpp  mandelbrot.cc)
add_subdirectory(tests)

target_compile_features(${PROJECT_NAME} PUBLIC cxx_std_20)
set (CMAKE_EXPORT_COMPILE_COMMANDS true)
