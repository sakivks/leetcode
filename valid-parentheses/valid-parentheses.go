package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "()[]"
	fmt.Println(isValid(input)) //true
}

func isValid(s string) bool {
	braces := strings.Split(s, "")
	bracketMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := []string{}

	for _, v := range braces {
		if val, present := bracketMap[v]; present {
			if len(stack) != 0 && stack[len(stack)-1] == val {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, v)
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
