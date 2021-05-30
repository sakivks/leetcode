package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums) //should print [1,3,12,0,0]
}

func moveZeroes(nums []int) {
	index := 0
	for k, val := range nums {
		if val != 0 {
			nums[index], nums[k] = nums[k], nums[index]
			index += 1
		}
	}
}
