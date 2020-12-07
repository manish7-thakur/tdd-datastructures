package word_search

//Source: https://leetcode.com/problems/word-search-ii/

func Exists(board [][]byte, word string, m, n, l, b int, visited [][]bool) bool {
	if len(word) == 0 {
		return true
	}
	visited[m][n] = true
	currByte := word[0]
	path1 := false
	path2 := false
	path3 := false
	path4 := false
	if m < l {
		rowidx := m + 1
		if !visited[rowidx][n] && currByte == board[rowidx][n] {
			path1 = Exists(board, word[1:], rowidx, n, l, b, visited)
			visited[rowidx][n] = false
		}
	}
	if m > 0 {
		rowidx := m - 1
		if !visited[rowidx][n] && currByte == board[rowidx][n] {
			path2 = Exists(board, word[1:], rowidx, n, l, b, visited)
			visited[rowidx][n] = false
		}
	}
	if n > 0 {
		colidx := n - 1
		if !visited[m][colidx] && currByte == board[m][colidx] {
			path3 = Exists(board, word[1:], m, colidx, l, b, visited)
			visited[m][colidx] = false
		}
	}
	if n < b {
		colidx := n + 1
		if !visited[m][colidx] && currByte == board[m][colidx] {
			path4 = Exists(board, word[1:], m, colidx, l, b, visited)
			visited[m][colidx] = false
		}
	}
	visited[m][n] = false
	return path1 || path2 || path3 || path4
}

func getAllIndices(row []byte, searchByte uint8) []int {
	indices := make([]int, 0)
	for i, byte := range row {
		if searchByte == byte {
			indices = append(indices, i)
		}
	}
	return indices
}

func FindWord(board [][]byte, word string, l, b int, visited [][]bool) bool {
	firstByte := word[0]
	for m, row := range board {
		idxs := getAllIndices(row, firstByte)
		for _, idx := range idxs {
			if idx != -1 && Exists(board, word[1:], m, idx, l, b, visited) {
				return true
			}
		}
	}
	return false
}

func FindWords(board [][]byte, words []string, l, b int) []string {
	foundWords := make([]string, 0)
	visited := make([][]bool, l+1)
	for i, _ := range visited {
		visited[i] = make([]bool, b+1)
	}
	for _, word := range words {
		if FindWord(board, word, l, b, visited) {
			foundWords = append(foundWords, word)
		}
	}
	return foundWords
}
