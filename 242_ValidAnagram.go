package main

import "fmt"

func isAnagram(s string, t string) bool {
    letterMap := make(map[rune]int)
	for _, letter := range s {
		if _, exists := letterMap[letter]; exists {
			letterMap[letter]++
		} else {
			letterMap[letter] = 1
		}
	}
	for _, letter := range t {
		if _, exits := letterMap[letter]; exits {
			letterMap[letter]--
			if letterMap[letter] < 0 {
				return false
			} else if letterMap[letter] == 0 {
				delete(letterMap, letter)
			}
		} else {
			return false
		}
	}
	if len(letterMap) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	s := "rat"
	t := "car"
	fmt.Println(isAnagram(s, t))
}