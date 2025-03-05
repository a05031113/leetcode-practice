package main

import "fmt"


func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	var backtrack func(iIndex int, jIndex int, wordIndex int) bool
	backtrack = func(iIndex int, jIndex int, wordIndex int) bool {
		if iIndex < 0 || iIndex >= len(board) || jIndex < 0 || jIndex >= len(board[0]) || word[wordIndex] != board[iIndex][jIndex] {
			return false
		}
		if wordIndex == len(word) - 1 {
			return true
		}
		tmp := board[iIndex][jIndex]
		board[iIndex][jIndex] = '#'
		found := backtrack(iIndex-1, jIndex, wordIndex+1) || 
			backtrack(iIndex+1, jIndex, wordIndex+1) || 
			backtrack(iIndex, jIndex-1, wordIndex+1) || 
			backtrack(iIndex, jIndex+1, wordIndex+1)
		board[iIndex][jIndex] = tmp
	return found
	}
    for i:=0; i<len(board); i++ {
		for j:=0; j<len(board[i]); j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}


func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word := "ABCCED"
	fmt.Println(exist(board, word))
}	