package main

import "fmt"

func main() {
	fmt.Println(climbStairs(6)) // should print 3
}

//Dymamic programming 
func climbStairs(n int) int {
	dp := []int{0,1,2}
	if n <= 2 {
		return dp[n]
	}

	for i := 3; i <= n; i++ {
		dp = append(dp,dp[i-1] + dp[i-2])
	}
	return dp[n]
}



// Recursive Approach
func climbStairs_recursive(n int) int {
	answerMap := map[int]int {
		1: 1,
		2: 2,
	}
	return climbStairsMemo(n, answerMap)
}

func climbStairsMemo(n int, answerMap map[int]int) int {
	if val, present := answerMap[n]; present {
		return val
	}
	answer := climbStairsMemo(n-1, answerMap) +climbStairsMemo(n-2, answerMap)
	answerMap[n] = answer
	return answer
}