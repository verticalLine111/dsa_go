package main

type Stack struct { 
	items []int
	size int
	capacity int
}

func (s *Stack) Push(value int)  {
	if s.size == s.capacity{ 
		s.ReSize()
	}
	s.items[s.size] = value
	s.size++
}

func (s *Stack) Pop() int {
	if s.size >0 {
		s.size--
		return s.items[s.size]
	} else {
		panic("stack is empty")
	}
}

func (s *Stack) Peek() int {
	if s.size >0 {
		return s.items[s.size-1]
	} else {
		panic("stack is empty")
	}
}

func (s *Stack) IsEmpty() bool {
	if s.size > 0 { 
		return false
	}
	return true
}

func (s *Stack) Size() int {
	return s.size
}
func (s *Stack) ReSize()  {
	newArr := make([]int , s.capacity * 2)
	for i := 0; i < s.size; i++ {
		newArr[i] = s.items[i]
	}
	s.items = newArr
	s.capacity *= 2
}

func main()  {
	stack := &Stack{
		size: 0,
		capacity: 4,
		items: make([]int, 4),
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	for !stack.IsEmpty() {
		println(stack.Pop())
	}
}