package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(5, 7)

	if total != 12 {
		t.Errorf("Sum was incorrect, got: %d, want: %d", total, 12)
	}
}

func TestDoMath(t *testing.T) {
	a := doMath(15, 5, Add)
	s := doMath(15, 5, Subtract)

	if a != 20 {
		log.Fatalf("Sum was incorrect, got: %d, want: %d", a, 20)
	}

	if s != 10 {
		t.Errorf("Difference was incorrect, got: %d, want: %d", s, 10)
	}
}

func TestFactorial(t *testing.T) {
	got := factorial(5)
	want := 120

	if got != want {
		log.Fatalf("Error: want: %v, got: %v", want, got)
	}
}
