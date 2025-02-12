package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
    buy := 0

	profit := 0

	for sold := 1; sold<len(prices); sold++ {
		if prices[sold] - prices[buy] > profit {
			profit = prices[sold] - prices[buy]
		}
		if prices[sold] < prices[buy] {
			buy = sold
		}
	} 
	return profit
}

func main() {
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(prices))
}