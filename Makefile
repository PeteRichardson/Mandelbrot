BUILD_DIR = build
BINARIES =	${BUILD_DIR}/mandelbrot-cpp   \
			${BUILD_DIR}/mandelbrot_tests \
			${BUILD_DIR}/mandelbrot-go	  \
			${BUILD_DIR}/mandelbrot-rust

all: $(BINARIES)

${BUILD_DIR}: 
	@mkdir -p ${BUILD_DIR} >> /dev/null

${BUILD_DIR}/mandelbrot-cpp: CPP_BUILD
	cp cpp/build/mandelbrot-cpp build/

${BUILD_DIR}/mandelbrot_tests: CPP_BUILD
	cp cpp/build/tests/mandelbrot_tests build/

CPP_CONFIGURE: CPP_BUILD_DIR
	@cmake -S cpp -B cpp/build

CPP_BUILD_DIR:
	@mkdir -p cpp/build >> /dev/null

CPP_BUILD: CPP_CONFIGURE cpp/*.cpp cpp/*.cc cpp/*.h
	@cmake --build cpp/build

CPP_TEST: ${BUILD_DIR}/mandelbrot_tests
	${BUILD_DIR}/mandelbrot_tests

${BUILD_DIR}/mandelbrot-go: GO_BUILD

GO_BUILD: go/main.go go/mandelbrot.go
	@cd go && go build -o ../${BUILD_DIR} && cd ..

GO_TEST: go/mandelbrot_test.go
	@cd go && go test -v . && cd ..

${BUILD_DIR}/mandelbrot-rust: RUST_BUILD
	@cp rust/target/debug/mandelbrot ${BUILD_DIR}/mandelbrot-rust

RUST_BUILD: rust/src/main.rs rust/src/mandelbrot.rs
	@cd rust && cargo build && cd ..

RUST_TEST: 
	@cd rust && cargo test && cd ..

test: CPP_TEST GO_TEST RUST_TEST

run: ${BINARIES}
	@${BUILD_DIR}/mandelbrot-rust mandelbrot-rust.png 1000x750 -1.20,0.35 -1,0.20
	@${BUILD_DIR}/mandelbrot-cpp
	@${BUILD_DIR}/mandelbrot-go
	@open *.png

clean: ${BUILD_DIR}
	rm -f ${BUILD_DIR}/mandelbrot*

fullclean: ${BUILD_DIR} clean
	rm -rf rust/target
	rm -rf cpp/build