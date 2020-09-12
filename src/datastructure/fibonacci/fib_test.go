package fibonacci

import "testing"

func TestFibonacciZeroNum(t *testing.T) {
	gen := New(0)
	n := gen.generateRec(0)
	expected := 0
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciFirstNum(t *testing.T) {
	gen := New(1)
	n := gen.generateRec(1)
	expected := 1
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciSecondNum(t *testing.T) {
	gen := New(2)
	n := gen.generateRec(2)
	expected := 1
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciThirdNum(t *testing.T) {
	gen := New(3)
	n := gen.generateRec(3)
	expected := 2
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciFiftiethNumRec(t *testing.T) {
	gen := New(50)
	n := gen.generateRec(50)
	expected := 12586269025
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciZeroNumIter(t *testing.T) {
	gen := New(0)
	n := gen.generateIter(0)
	expected := 0
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciFirstNumIter(t *testing.T) {
	gen := New(0)
	n := gen.generateIter(1)
	expected := 1
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciSecondNumIter(t *testing.T) {
	gen := New(0)
	n := gen.generateIter(2)
	expected := 1
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciThirdNumIter(t *testing.T) {
	gen := New(0)
	n := gen.generateIter(3)
	expected := 2
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciFourthNumIter(t *testing.T) {
	gen := New(0)
	n := gen.generateIter(4)
	expected := 3
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciFifthNumIter(t *testing.T) {
	gen := New(0)
	n := gen.generateIter(5)
	expected := 5
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func TestFibonacciFiftiethNumIter(t *testing.T) {
	gen := New(1)
	n := gen.generateIter(50)
	expected := 12586269025
	if n != expected {
		t.Errorf("Expected %d but found %d", expected, n)
	}
}

func BenchmarkFibonacciHundredthSum(b *testing.B) {
	n := 100
	gen := New(n)
	for i := 0; i < b.N; i++ {
		gen.generateRec(n)
	}
}
