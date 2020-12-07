package word_search

import (
	"testing"
)

func TestSearchEmptyWord(t *testing.T) {
	baord := [][]byte{{'a'}}
	visited := make([][]bool, 1)
	visited[0] = make([]bool, 1)
	actual := Exists(baord, "", 0, 0, 0, 0, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchSingleCharExistentWord(t *testing.T) {
	board := [][]byte{{'a'}}
	word := "a"
	visited := make([][]bool, 1)
	visited[0] = make([]bool, 1)
	actual := Exists(board, word[1:], 0, 0, 0, 0, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharNonExistentWord(t *testing.T) {
	board := [][]byte{{'a'}}
	word := "ab"
	visited := make([][]bool, 1)
	visited[0] = make([]bool, 1)
	actual := Exists(board, word[1:], 0, 0, 0, 0, visited)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharExistentWordDownDirection(t *testing.T) {
	board := [][]byte{{'a'}, {'b'}}
	word := "ab"
	visited := make([][]bool, 2)
	visited[0] = make([]bool, 1)
	visited[1] = make([]bool, 1)
	actual := Exists(board, word[1:], 0, 0, 2, 1, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharExistentWordUpDirection(t *testing.T) {
	board := [][]byte{{'a'}, {'b'}}
	word := "ba"
	visited := make([][]bool, 2)
	visited[0] = make([]bool, 1)
	visited[1] = make([]bool, 1)
	actual := Exists(board, word[1:], 1, 0, 1, 0, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharExistentWordBackwardDirection(t *testing.T) {
	board := [][]byte{{'a', 'b'}}
	word := "ba"
	visited := make([][]bool, 1)
	visited[0] = make([]bool, 2)
	actual := Exists(board, word[1:], 0, 1, 0, 1, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharExistentWordForwardDirection(t *testing.T) {
	board := [][]byte{{'a', 'b'}}
	word := "ab"
	visited := make([][]bool, 1)
	visited[0] = make([]bool, 2)
	actual := Exists(board, word[1:], 0, 0, 1, 2, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchTripleCharNonExistentWord(t *testing.T) {
	board := [][]byte{{'a', 'b', 'c'}}
	word := "abx"
	visited := make([][]bool, 1)
	visited[0] = make([]bool, 3)
	actual := Exists(board, word[1:], 0, 0, 0, 2, visited)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchTripleCharExistentWord(t *testing.T) {
	board := [][]byte{{'a', 'b'}, {'x', 'c'}}
	word := "abc"
	visited := make([][]bool, 2)
	visited[0] = make([]bool, 2)
	visited[1] = make([]bool, 2)
	actual := Exists(board, word[1:], 0, 0, 2, 2, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchTripleCharExistentWordTwoDirections(t *testing.T) {
	board := [][]byte{
		{'a', 'b', 'x'},
		{'b', 'v', 'd'},
		{'x', 'c', 'b'}}
	word := "abxd"
	visited := make([][]bool, 3)
	visited[0] = make([]bool, 3)
	visited[1] = make([]bool, 3)
	visited[2] = make([]bool, 3)
	actual := Exists(board, word[1:], 0, 0, 3, 3, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchTripleCharExistentWordTwoDirectionsVisted(t *testing.T) {
	board := [][]byte{
		{'a', 'b', 'c'},
		{'a', 'e', 'd'},
		{'a', 'b', 'g'}}
	word := "eaabcdgba"
	visited := make([][]bool, 3)
	visited[0] = make([]bool, 3)
	visited[1] = make([]bool, 3)
	visited[2] = make([]bool, 3)
	actual := Exists(board, word[1:], 1, 1, 3, 3, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchTripleCharNonExistentWordAllDirectionsVisitedCell(t *testing.T) {
	board := [][]byte{{'a', 'b'}, {'x', 'c'}}
	word := "aba"
	visited := make([][]bool, 2)
	visited[0] = make([]bool, 2)
	visited[1] = make([]bool, 2)
	actual := Exists(board, word[1:], 0, 0, 1, 1, visited)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	word = "bab"
	actual = Exists(board, word[1:], 0, 1, 1, 1, visited)
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	word = "cbc"
	actual = Exists(board, word[1:], 1, 1, 1, 1, visited)
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	word = "bcb"
	actual = Exists(board, word[1:], 0, 1, 1, 1, visited)
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchFourCharNonExistentWord(t *testing.T) {
	board := [][]byte{{'a', 'b'}, {'c', 'd'}}
	word := "abcb"
	visited := make([][]bool, 2)
	visited[0] = make([]bool, 2)
	visited[1] = make([]bool, 2)
	actual := Exists(board, word[1:], 0, 0, 1, 1, visited)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchFourCharExistentWordBiggerBoard(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'}}
	word := "eat"
	visited := make([][]bool, 4)
	visited[0] = make([]bool, 4)
	visited[1] = make([]bool, 4)
	visited[2] = make([]bool, 4)
	visited[3] = make([]bool, 4)
	actual := Exists(board, word[1:], 1, 3, 3, 3, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = Exists(board, word[1:], 1, 0, 3, 3, visited)
	expected = false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	word = "oath"
	actual = Exists(board, word[1:], 0, 0, 3, 3, visited)
	expected = true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestFindSingleCharWordNonExistentFirstChar(t *testing.T) {
	board := [][]byte{{'w'}}
	visited := make([][]bool, 1)
	visited[0] = make([]bool, 1)
	actual := FindWord(board, "x", 0, 0, visited)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestFindSingleCharWordExistentFirstChar(t *testing.T) {
	board := [][]byte{{'w'}}
	visited := make([][]bool, 1)
	visited[0] = make([]bool, 1)
	actual := FindWord(board, "w", 0, 0, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestFindDoubleCharWordExistentFirstCharMultipleOccurrence(t *testing.T) {
	board := [][]byte{{'w', 'x'}, {'w', 'y'}}
	visited := make([][]bool, 2)
	visited[0] = make([]bool, 2)
	visited[1] = make([]bool, 2)
	actual := FindWord(board, "wy", 2, 2, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestFindDoubleCharWordExistentFirstCharMultipleOccurrenceSameRow(t *testing.T) {
	board := [][]byte{{'w', 'w'}, {'x', 'y'}}
	visited := make([][]bool, 2)
	visited[0] = make([]bool, 2)
	visited[1] = make([]bool, 2)
	actual := FindWord(board, "wy", 2, 2, visited)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestFindWordsEmptyWords(t *testing.T) {
	visited := make([][]bool, 1)
	visited[0] = make([]bool, 1)
	res := FindWords([][]byte{{'w'}}, []string{}, 1, 1)
	if len(res) != 0 {
		t.Errorf("expected empty but found %v", res)
	}
}

func TestFindWordsExistentSingleWord(t *testing.T) {
	visited := make([][]bool, 1)
	visited[0] = make([]bool, 1)
	res := FindWords([][]byte{{'w'}}, []string{"w"}, 1, 1)
	actual := [1]string{}
	copy(actual[:], res)
	expected := [1]string{"w"}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestFindWordsExistentMultipleWords(t *testing.T) {
	res := FindWords([][]byte{{'w', 'x'}, {'x', 'y'}}, []string{"wx", "cy", "xy"}, 2, 2)
	actual := [3]string{}
	copy(actual[:], res)
	expected := [3]string{"wx", "xy"}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}
