package main

import "fmt"

var profit int = 0

func maxProfit(prices []int) int {
	for day, price := range prices {
		nextDaysCompute(prices, day, price)
	}
	return profit
}

func nextDaysCompute(prices []int, curDay, curPrice int) int {
	for day, price := range prices {
		if day > curDay && price > curPrice {
			totalProfit := price - curPrice
			if totalProfit > profit {
				profit = totalProfit
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	profit = 0
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
