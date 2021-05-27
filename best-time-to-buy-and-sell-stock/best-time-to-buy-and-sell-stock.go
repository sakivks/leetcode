package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices)) // should print 5
}

func maxProfit(prices []int) int {
	minPrice, maxPro := prices[0], 0

	for _, val := range prices {
		if val < minPrice {
			minPrice = val
		}
		profit := val - minPrice
		if profit > maxPro {
			maxPro = profit
		}
	}

	return maxPro
}

//Using brute force
func maxProfit_brute(prices []int) int {
	maxP := 0
	for i, v1 := range prices {
		for _, v2 := range prices[i:] {
			if v2-v1 > maxP {
				maxP = v2 - v1
			}
		}
	}
	return maxP
}
