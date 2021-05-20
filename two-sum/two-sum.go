package main

import "fmt"

func main() {
	nums, target := []int{2, 7, 11, 15}, 9
	fmt.Println(twoSum(nums, target))
}

// hash table
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for k, v := range nums {
		if val, present := hashMap[target-v]; present {
			return []int{k, val}
		}
		hashMap[v] = k
	}
	return nums
}

// brute force O(n^2)
func twoSumBrute(nums []int, target int) []int {
	for i, v1 := range nums {
		for j, v2 := range nums[i+1:] {
			if v1+v2 == target {
				return []int{i, j}
			}
		}
	}
	return nums
}
