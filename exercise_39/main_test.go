package main

import (
	"testing"
)

func TestIsEven(t *testing.T) {
	trueCase := 2
	falseCase := 3
	if isEven(trueCase) != true {
		t.Errorf("Expected %v to be even", trueCase)
	}
	if isEven(falseCase) != false {
		t.Errorf("Expected %v to be odd", falseCase)
	}
}
