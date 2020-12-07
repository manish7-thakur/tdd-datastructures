package word_search

//Source: https://leetcode.com/problems/word-search-ii/

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

func getAllIndices(row []byte, searchByte uint8) []int {
	indices := make([]int, 0)
	for i, byte := range row {
		if searchByte == byte {
			indices = append(indices, i)
		}
	}
	return indices
}

func Exists(board [][]byte, word string, m, n, l, b int, visited [][]bool) bool {
	if len(word) == 0 {
		return true
	}
	visited[m][n] = true
	currByte := word[0]
	if m+1 < l && !visited[m+1][n] && currByte == board[m+1][n] {
		if Exists(board, word[1:], m+1, n, l, b, visited) {
			visited[m][n] = false
			return true // if found, return from here & no need to check below statements, `return Exists()` may be false & prevent traversing below paths
		}
	}
	if m-1 >= 0 && !visited[m-1][n] && currByte == board[m-1][n] {
		if Exists(board, word[1:], m-1, n, l, b, visited) {
			visited[m][n] = false
			return true
		}
	}
	if n-1 >= 0 && !visited[m][n-1] && currByte == board[m][n-1] {
		if Exists(board, word[1:], m, n-1, l, b, visited) {
			visited[m][n] = false
			return true
		}
	}
	if n+1 < b && !visited[m][n+1] && currByte == board[m][n+1] {
		if Exists(board, word[1:], m, n+1, l, b, visited) {
			visited[m][n] = false
			return true
		}
	}
	visited[m][n] = false
	return false
}
