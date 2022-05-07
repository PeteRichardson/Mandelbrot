// hello_test.go - This is the unit test file
package main

import (
	"testing"
)

func TestEscapeTime(t *testing.T) {
	actual := escape_time(complex(0, 0), 255)
	if actual != 255 {
		t.Errorf("expected '1', got '%d'", actual)
	}
	actual = escape_time(complex(0.5, 0.3), 255)
	if actual != 6 {
		t.Errorf("expected '6', got '%d'", actual)
	}
	actual = escape_time(complex(-0.5, 0.25), 255)
	if actual != 255 {
		t.Errorf("expected '0', got '%d'", actual)
	}
	actual = escape_time(complex(0.0, 0.75), 255)
	if actual != 34 {
		t.Errorf("expected '33', got '%d'", actual)
	}
}

func TestPixelToPoint(t *testing.T) {
	actual := pixel_to_point(100, 200, 25, 175, complex(-1.0, 1.0), complex(1.0, -1.0))
	// if actual != complex(-0.5, -0.75) {
	if actual != complex(-0.5, -0.75) {
		t.Errorf("expected complex(-0.5, -0.75), got '%v'", actual)
	}
}
