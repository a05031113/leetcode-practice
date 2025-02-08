package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	max := 0
    
	for right > left {
		floor := right - left

		var h int
		if height[right] < height[left] {
			h = height[right]
			right --
		} else {
			h = height[left]
			left ++
		}
		if h * floor > max {
			max = h * floor
		}
	}
	return max
}


func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}