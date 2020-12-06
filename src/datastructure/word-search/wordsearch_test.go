package word_search

import "testing"

func TestSearchEmptyWord(t *testing.T) {
	baord := [][]byte{{'a'}}
	actual := Exists(baord, "", 0, 0, 0, 0)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchSingleCharExistentWord(t *testing.T) {
	board := [][]byte{{'a'}}
	word := "a"
	actual := Exists(board, word[1:], 0, 0, 0, 0)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharNonExistentWord(t *testing.T) {
	board := [][]byte{{'a'}}
	word := "ab"
	actual := Exists(board, word[1:], 0, 0, 0, 0)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharExistentWordDownDirection(t *testing.T) {
	board := [][]byte{{'a'}, {'b'}}
	word := "ab"
	actual := Exists(board, word[1:], 0, 0, 1, 0)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharExistentWordUpDirection(t *testing.T) {
	board := [][]byte{{'a'}, {'b'}}
	word := "ba"
	actual := Exists(board, word[1:], 1, 0, 1, 0)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharExistentWordBackwardDirection(t *testing.T) {
	board := [][]byte{{'a', 'b'}}
	word := "ba"
	actual := Exists(board, word[1:], 0, 1, 0, 1)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharExistentWordForwardDirection(t *testing.T) {
	board := [][]byte{{'a', 'b'}}
	word := "ab"
	actual := Exists(board, word[1:], 0, 0, 0, 1)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchTripleCharNonExistentWord(t *testing.T) {
	board := [][]byte{{'a', 'b', 'c'}}
	word := "abx"
	actual := Exists(board, word[1:], 0, 0, 0, 2)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchTripleCharExistentWord(t *testing.T) {
	board := [][]byte{{'a', 'b'}, {'x', 'c'}}
	word := "abc"
	actual := Exists(board, word[1:], 0, 0, 1, 1)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func Exists(board [][]byte, word string, m, n, l, b int) bool {
	if len(word) == 0 {
		return true
	}
	currByte := word[0]
	if m < l {
		if currByte == board[m+1][n] {
			return Exists(board, word[1:], m+1, n, l, b)
		}
	}
	if m > 0 {
		if currByte == board[m-1][n] {
			return Exists(board, word[1:], m-1, n, l, b)
		}
	}
	if n > 0 {
		if currByte == board[m][n-1] {
			return Exists(board, word[1:], m, n-1, l, b)
		}
	}
	if n < b {
		if currByte == board[m][n+1] {
			return Exists(board, word[1:], m, n+1, l, b)
		}
	}
	return false
}
