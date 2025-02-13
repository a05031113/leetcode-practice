package main

import "fmt"

func isValid(s string) bool {
    var stack []string
	for i:=0; i<len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			stack = append(stack, string(s[i]))
			continue
		} 
        if len(stack) == 0 {
            return false
        }
		if string(s[i]) == ")" && stack[len(stack)-1] != "(" {
			return false
		} else if string(s[i]) == "]" && stack[len(stack)-1] != "[" {
			return false
		} else if string(s[i]) == "}" && stack[len(stack)-1] != "{" {
			return false
		} else {
			stack = append(stack[:len(stack)-1])
		}
	}
	if len(stack) != 0 {
		return false
	} 
	return true
}

func main() {
	s := "()"
	fmt.Println(isValid(s))
}