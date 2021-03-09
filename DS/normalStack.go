package main

import (
	"fmt"
)

type Stack struct {
	pool []string
}

// Check if the stack is empty and return appropriately
func (s *Stack) UnderFlow() bool {
	return len(s.pool) == 0
}

// Push the value onto the stack
func (s *Stack) Push(v string) {
	s.pool = append(s.pool, v)
}

// Pop values from the stack
func (s *Stack) Pop() (bool, string) {

	if s.UnderFlow() {
		return true, ""
	} else {
		index := len(s.pool) - 1 // Get the index of the top most element.
		removed := s.pool[index] // Index into the slice and obtain the element.
		s.pool = s.pool[:index] // Remove it from the stack by slicing it off.
		return true, removed
	}
}

func main() {
	stack := Stack{}
	fmt.Println(stack)
	stack.Push("100")
	stack.Push("200")
	stack.Push("300")
	fmt.Println(stack)

	for len(stack.pool) > 0 {
		check, value := stack.Pop()
		if check == true {
			fmt.Println(value)
		}
	}
}