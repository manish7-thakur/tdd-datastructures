package sparsematmul

import "testing"

func TestMulOneCrossOneMat(t *testing.T) {
	a := [3][3]int{{1, 0, 0},{1, 0, 0},{0, 2, 3}}
	b := [3][3]int{{0, 0, 1}, {1, 0, 0}, {0, 1, 0}}

	actual := multiply(a, b)
	expected := [3][3]int{{0,0,1}, {0,0,1}, {2,3,0}}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual);
	}
}