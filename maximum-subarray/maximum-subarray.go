package main

import "fmt"

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(input)) // should print 6
}

func maxSubArray(nums []int) int {
	sum, maxSum := 0, nums[0]

	for _, val := range nums {
		sum += val
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}
