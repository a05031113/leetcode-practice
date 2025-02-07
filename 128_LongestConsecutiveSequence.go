package main

import "fmt"


func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})

	for _, item := range nums {
		set[item] = struct{}{}
	}

	maxCount := 0
    for num := range set {
        if _, exists := set[num-1]; !exists {
            curr := num
            currStreak := 1
            
            for {
                curr++ 
                if _, exists := set[curr]; !exists {
                    break 
                }
                currStreak++
            }
            
            if currStreak > maxCount {
                maxCount = currStreak
            }
        }
    }
	return maxCount
}

func main() {
	nums := []int{100,4,200,1,3,2} 
	fmt.Println(longestConsecutive(nums))
}