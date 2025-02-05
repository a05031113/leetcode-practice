package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
    r := []rune(s)
    sort.Slice(r, func(i, j int) bool {
        return r[i] < r[j]
    })
    return string(r)
}


func groupAnagrams(strs []string) [][]string {
	wordMap := make(map[string][]string)
    for _, str := range strs {
		sortResult := sortString(str)
		wordMap[sortResult] = append(wordMap[sortResult], str)
	}
	result := make([][]string, 0, len(wordMap))
    for _, group := range wordMap {
        result = append(result, group)
    }
	fmt.Println(wordMap)
	return result
}

func main() {
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams(strs))
}