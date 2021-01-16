package sparsematmul

import (
	"testing"
	"math/rand"
	"time"
)

func TestMulOneCrossOneMat(t *testing.T) {
	a := [3][3]int{{1, 0, 0},{1, 0, 0},{0, 2, 3}}
	b := [3][3]int{{0, 0, 1}, {1, 0, 0}, {0, 1, 0}}

	actual := multiply(a, b)
	expected := [3][3]int{{0,0,1}, {0,0,1}, {2,3,0}}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual);
	}
	actual = multiply2(a, b)
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual);
	}
}

func BenchmarkMultiply(b *testing.B) {
	mat1 := randomSparseMat()
	mat2 := randomSparseMat()
	for i :=0; i < b.N ; i++ {
		multiply(mat1, mat2)
	}
}

func BenchmarkMultiply2(b *testing.B) {
	mat1 := randomSparseMat()
	mat2 := randomSparseMat()
	for i :=0; i < b.N ; i++ {
		multiply2(mat1, mat2)
	}
}

func randomSparseMat() [3][3]int {
	rand.Seed(time.Now().UnixNano())
	mat := [3][3]int{};
	for i := 0; i< len(mat); i++ {
		for j :=0; j < len(mat); j++ {
			if(rand.Intn(2) == 0) {
				mat[i][j] = rand.Intn(10);
			}
		}
	}
	return mat;
}