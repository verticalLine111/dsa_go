package main

import (
	"fmt"
)

type Stack struct {
	elements []int // C의 int* array 역할
	top      int   // 현재 데이터 개수 (다음 데이터가 들어갈 위치)
	capacity int   // 현재 할당된 메모리 크기
}

// 초기 용량을 지정하여 스택 생성
func NewStack(initialCapacity int) *Stack {
	return &Stack{
		elements: make([]int, initialCapacity),
		top:      0,
		capacity: initialCapacity,
	}
}

func (s *Stack) Push(value int) {
	// 1. 용량이 꽉 찼는지 확인 (C의 realloc 준비)
	if s.top == s.capacity {
		s.grow()
	}
	// 2. 값 추가 및 인덱스 이동
	s.elements[s.top] = value
	s.top++
}

// 수동 메모리 재할당 함수 (append 내부 로직 모사)
func (s *Stack) grow() {
	// 새로운 용량 결정 (보통 2배씩 확장하여 Amortized O(1) 달성)
	newCapacity := s.capacity * 2
	if newCapacity == 0 {
		newCapacity = 1
	}

	// 새로운 메모리 공간 할당 (C의 malloc)
	newElements := make([]int, newCapacity)

	// 기존 데이터 복사 (C의 memcpy)
	for i := 0; i < s.top; i++ {
		newElements[i] = s.elements[i]
	}

	// 새 배열로 교체 및 용량 업데이트
	s.elements = newElements
	s.capacity = newCapacity
	fmt.Printf("--- Memory Resized: %d -> %d ---\n", s.capacity/2, s.capacity)
}

func (s *Stack) Pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	// 인덱스를 먼저 줄이고 값을 가져옴
	s.top--
	value := s.elements[s.top]
	// 선택: 메모리 절약을 위해 여기서 0으로 초기화할 수도 있음
	return value, true
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func main() {
	// 초기 용량 2로 시작 (재할당 과정을 보기 위함)
	stack := NewStack(2)

	fmt.Println("Pushing: 10, 20, 30, 40, 50")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30) // 여기서 첫 번째 재할당 발생 (2->4)
	stack.Push(40)
	stack.Push(50) // 여기서 두 번째 재할당 발생 (4->8)

	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Println("Popped:", value)
	}
}
