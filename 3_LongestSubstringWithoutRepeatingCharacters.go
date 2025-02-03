package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := 0
	lastIndex := 0
	tmp := make(map[rune]int)
	
    for i, val := range s {
		if index, exists := tmp[val]; exists && index >= lastIndex {
			lastIndex = index + 1
		}
		currentLen := i - lastIndex + 1
		if currentLen > max {
            max = currentLen
        }
		tmp[val] = i
	}

	return max
}

func main() {
	s := "abba"
	fmt.Println(lengthOfLongestSubstring(s))
}