package main

import "fmt"

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums)) //should print 4
}

//bitwise XOR, every bits which appears twice get's cancelled out
func singleNumber(nums []int) int {
	a := 0
	for _, val := range nums {
		a ^= val
	}
	return a
}
