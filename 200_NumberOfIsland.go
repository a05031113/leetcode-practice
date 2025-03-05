package main

import "fmt"


func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
		return 0
	}
	count := 0

	var backtrack func(iIndex int, jIndex int) bool
	backtrack = func(iIndex int, jIndex int) bool {
		if iIndex < 0 || iIndex >= len(grid) || jIndex < 0 || jIndex >= len(grid[0]) || grid[iIndex][jIndex] != '1' {
			return false
		}

		// tmp := grid[iIndex][jIndex]
		grid[iIndex][jIndex] = '#'
		backtrack(iIndex-1, jIndex)
		backtrack(iIndex+1, jIndex)
		backtrack(iIndex, jIndex-1)
		backtrack(iIndex, jIndex+1)
		return true
	}
    for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[i]); j++ {
			if backtrack(i, j) {
				count ++
			}
		}
	}
	return count

}


func main() {
	grid := [][]byte{
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'},
	}

	fmt.Println(numIslands(grid))
}