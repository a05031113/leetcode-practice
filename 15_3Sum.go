package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
            continue
        }

		left, right := i+1, len(nums)-1
		target := -nums[i]

		for right > left {
			sum := nums[left] + nums[right]

			if target == sum {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left + 1] {
					left ++
				}
				for left < right && nums[right] == nums[right - 1] {
					right --
				}
				left ++
				right --
			} else if target > sum {
				left ++
			} else {
				right --
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1,0,1,2,-1,-4}

	fmt.Println(threeSum(nums))
}