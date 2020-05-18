package main

import (
	"testing"
)


func TestAdd(t *testing.T) {
	result := Add(4,5)
	if result != 9 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 3)
	}
}
func TestSubtract(t *testing.T) {
	result := Subtract(5,3)
	if result != 2 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 2)
	}
}
func TestMultiply(t *testing.T) {
	result := Multiply(4,5)
	if result != 20 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 20)
	}
}
func TestDivision(t *testing.T) {
	result := Division(10,2)
	if result != 5 {
		t.Errorf("Result was incorrect, got: %f, want: %f.", result, float64(2))
	}
}