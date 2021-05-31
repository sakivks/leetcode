package main

import "fmt"

func main() {
	fmt.Println(countBits(5)) // should return [0,1,1,2,1,2]
}

func countBits(n int) []int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = dp[i/2]
		if i%2 == 1 {
			dp[i] += 1
		}
	}
	return dp
}
