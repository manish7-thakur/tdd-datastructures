package fibonacci

import "testing"

func TestFibonacciZeroNum(t *testing.T) {
	gen := Fib{}
	n := gen.generate(0)
	expected := 0
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciFirstNum(t *testing.T) {
	gen := Fib{}
	n := gen.generate(1)
	expected := 1
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciSecondNum(t *testing.T) {
	gen := Fib{}
	n := gen.generate(2)
	expected := 1
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciThirdNum(t *testing.T) {
	gen := Fib{}
	n := gen.generate(3)
	expected := 2
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciFiftiethNum(t *testing.T) {
	gen := Fib{}
	n := gen.generate(40)
	expected := 102334155
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}
