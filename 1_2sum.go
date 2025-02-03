package main

import "fmt"

func twoSum(nums []int, target int) []int {
    pairingMap := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        complement := target - nums[i]
        value, exists := pairingMap[complement]
        if exists {
            return []int{value, i}
        }
        pairingMap[nums[i]] = i
    }
    return []int{}
}

