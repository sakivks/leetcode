package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	fmt.Println(obj.Top())
}

type intStack struct {
	s []int
}

func (s *intStack) empty() bool {
	return len(s.s) == 0
}

func (s *intStack) push(n int) {
	s.s = append(s.s, n)
}

func (s *intStack) peek() int {
	if s.empty() {
		panic("stack is empty")
	}

	return s.s[len(s.s)-1]
}

func (s *intStack) pop() int {
	if s.empty() {
		panic("stack is empty")
	}

	val := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]

	return val
}

type MinStack struct {
	mins intStack
	nums intStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.nums.push(x)

	// we need to check if the mins stack is empty
	// the element has to be the smallest so far
	if this.mins.empty() || x <= this.mins.peek() {
		this.mins.push(x)
	}
}

func (this *MinStack) Pop() {
	val := this.nums.pop()

	if val == this.mins.peek() {
		this.mins.pop()
	}
}

func (this *MinStack) Top() int {
	return this.nums.peek()
}

func (this *MinStack) GetMin() int {
	return this.mins.peek()
}

//Old implementation
type MinStack_basic struct {
	min int
	arr []int
}

/** initialize your data structure here. */
func Constructor_old() MinStack_basic {
	var minStack MinStack_basic
	return minStack
}

func (this *MinStack_basic) Push(val int) {
	if len(this.arr) == 0 || this.min > val {
		this.min = val
	}
	this.arr = append(this.arr, val)
}

func (this *MinStack_basic) Pop() {
	if this.min == this.arr[len(this.arr)-1] {
		min := this.arr[0]
		for _, v := range this.arr[:len(this.arr)-1] {
			if v < min {
				min = v
			}
		}
		this.min = min
	}
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack_basic) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack_basic) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 */
