package fibonacci

import "testing"

func TestFibonacciZeroNum(t *testing.T) {
	gen := New(0)
	n := gen.generate(0)
	expected := 0
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciFirstNum(t *testing.T) {
	gen := New(1)
	n := gen.generate(1)
	expected := 1
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciSecondNum(t *testing.T) {
	gen := New(2)
	n := gen.generate(2)
	expected := 1
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciThirdNum(t *testing.T) {
	gen := New(3)
	n := gen.generate(3)
	expected := 2
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciFiftiethNum(t *testing.T) {
	gen := New(50)
	n := gen.generate(50)
	expected := 12586269025
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}
