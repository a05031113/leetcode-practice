package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
    var output [][]int
	var current []int

	var backtrack func(start int, remaining int)
	backtrack = func(start int, remaining int) {
		if remaining == 0 {
			combination := make([]int, len(current))
            copy(combination, current)
			output = append(output, combination)
			return
		}
		if remaining < 0 {
			return
		}
		for i:=start; i<len(candidates); i++ {
			current = append(current, candidates[i])
			backtrack(i, remaining-candidates[i])
			current = current[:len(current)-1]
		}
	}
	backtrack(0, target)
	return output
}


func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}