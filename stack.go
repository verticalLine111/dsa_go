package main

import (
	"fmt"
)
const defaultSize int = 1
type Stack struct {
	elements []int
	top int
	capacity int
}
func NewStack() *Stack {
	return &Stack{ 
		elements: make([]int, defaultSize),
		top: 0,
		capacity: defaultSize,
	}
}
func (s *Stack) Push(value int) {
	if s.top == s.capacity {
		s.resize()
	}
	s.elements[s.top] = value
	s.top++
}
func (s *Stack) resize(){
	newCapacity := s.capacity * 2
	if newCapacity == 0 {
		newCapacity = 1
	}

	newElements := make([]int, newCapacity)
	for i:=0; i <s.top ; i++ {
		newElements[i] = s.elements[i]
	}
	s.elements = newElements
	s.capacity = newCapacity
	fmt.Printf("-- Memory resized %d --> %d \n" , s.capacity/2 , s.capacity)
}

func (s *Stack) Pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	s.top--;
	value := s.elements[s.top]
	return value, true
}

func (s *Stack) isEmpty() bool {
	return s.top == 0;
}

func main() {
	s := NewStack()
	s.Push(2)
	s.Push(5)
	if value, ok := s.Pop(); ok {
		fmt.Printf("%d pop\n", value)
	}
	
}