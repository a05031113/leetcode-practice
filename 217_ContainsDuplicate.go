package main

import "fmt"

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool)
	for _, val := range nums {
		if _, exists := numsMap[val]; exists {
			return true
		}
		numsMap[val] = true
	}
    return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}