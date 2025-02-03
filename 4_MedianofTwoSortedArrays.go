package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	x, y := len(nums1), len(nums2)
    low, high := 0, x

    for low <= high {
        partitionX := (low + high) / 2
        partitionY := (x+y+1)/2 - partitionX

        // Find the elements around partitions
        var leftX, rightX, leftY, rightY int

        // Handle edge cases for X partition
        if partitionX == 0 {
            leftX = math.MinInt32
        } else {
            leftX = nums1[partitionX-1]
        }
        if partitionX == x {
            rightX = math.MaxInt32
        } else {
            rightX = nums1[partitionX]
        }
        // Handle edge cases for Y partition
        if partitionY == 0 {
            leftY = math.MinInt32
        } else {
            leftY = nums2[partitionY-1]
        }
        if partitionY == y {
            rightY = math.MaxInt32
        } else {
            rightY = nums2[partitionY]
        }
		fmt.Println(leftX, rightX, leftY, rightY)
        // Check if we found the right partition
        if leftX <= rightY && leftY <= rightX {
            // If total length is odd
            if (x+y)%2 == 1 {
                return float64(max(leftX, leftY))
            }
            // If total length is even
            return float64(max(leftX, leftY)+min(rightX, rightY)) / 2.0
        } else if leftX > rightY {
            // Need to move partition to left in X
            high = partitionX - 1
        } else {
            // Need to move partition to right in X
            low = partitionX + 1
        }
    }

    // Should never reach here if input arrays are sorted
    return 0.0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Helper function to find minimum of two integers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}