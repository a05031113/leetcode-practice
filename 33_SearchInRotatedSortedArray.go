package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left != right {
		index := int((left + right) / 2)

		if nums[index] == target {
			return index
		} else if nums[left] == target {
			return left
		} else if nums[right] == target {
			return right
		}
		if nums[index] > nums[right] {
			if nums[left] < target && target < nums[index] {
				right = index
			} else {
				left = index
			}
		} else if nums[index] < nums[right] {
			if nums[index] < target && target < nums[right] {
				left = index
			} else {
				right = index
			}
		}
		
		if right - 1 == left {
			break
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func main() {
	nums := []int{4,5,6,7,8,1,2,3}
	target := 8
	fmt.Println(search(nums, target))
}