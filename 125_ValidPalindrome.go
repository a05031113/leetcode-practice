package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
    left := 0
	right := len(s) - 1

	for right > left {
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
			left ++
			continue
		}
		if !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
			right--
			continue
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		right--
		left++
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"

	fmt.Println(isPalindrome(s))
}