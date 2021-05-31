package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})) //should print [5,6]
}

func findDisappearedNumbers(nums []int) []int {
	missing := []int{}
	for _, v := range nums {
		if v < 0 {
			v *= -1
		}
		if nums[v-1] > 0 {
			nums[v-1] *= -1
		}
	}
	for k, v := range nums {
		if v > 0 {
			missing = append(missing, k+1)
		}
	}
	return missing
}
