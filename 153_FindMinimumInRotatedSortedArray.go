package main

import "fmt"

func findMin(nums []int) int {
    left := 0
	right := len(nums) - 1

	for left != right {
		index := int((left + right) / 2)

		if nums[index] > nums[right] {
			left = index
		} else {
			right = index
		}
		if right - 1 == left {
			break
		}
	}
	output := min(nums[left], nums[right])
	return output
}

func main() {
	nums := []int{11,13,15,17}
	fmt.Println(findMin(nums))
}