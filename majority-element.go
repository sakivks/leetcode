package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums)) //should print 2
}

//using space O(1)
func majorityElement(nums []int) int {
	count, majorityEl := 0, nums[0]

	for k, v := range nums {
		if v == majorityEl {
			count++
		} else {
			count--
		}

		if count == 0 {
			majorityEl = nums[k+1]
		}
	}

	return majorityEl
}

//using hash map
func majorityElement_v1(nums []int) int {
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
