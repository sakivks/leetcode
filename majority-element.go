package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums)) //should print 2
}

func majorityElement(nums []int) int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v] += 1
	}

	majorityEl, majorityCount := nums[0], 0

	for k, v := range count {
		if v > majorityCount {
			majorityEl = k
			majorityCount = v
		}
	}
	return majorityEl
}
