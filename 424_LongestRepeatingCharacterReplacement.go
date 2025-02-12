package main

import "fmt"



func characterReplacement(s string, k int) int {
    runes := []rune(s)
	if len(runes) == 0 {
        return 0
    }
	
	charMap := make(map[rune]int)
	slow := 0
	maxLen := 0
	maxCount := 0
	for fast:=0; fast<len(runes); fast++ {
		charMap[runes[fast]] ++
		maxCount = max(maxCount, charMap[runes[fast]])

		for fast-slow+1 - maxCount > k {
            charMap[runes[slow]]--
            slow++
        }
		maxLen = max(maxLen, fast-slow+1)
	}
	return maxLen
}

func main() {
	s := "ABAB"
	k := 2
	fmt.Println(characterReplacement(s, k))
}