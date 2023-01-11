package Alg_for_array

import "fmt"

type stack struct {
	data int
	next *stack // ตัวเองเป็น pointer ไปยังตัวเอง
}

func (s *stack) push(data int) {
	node := &stack{data: data, next: s.next}
	s.next = node
}

func (s *stack) front() int {
	if s.next == nil {
		panic("Stack is empty")
	}
	return s.next.data
}

func (s *stack) pop() {
	if s.next == nil {
		panic("Stack is empty")
	}
	s.next = s.next.next
}

func (s *stack) empty() bool {
	return s.next == nil
}

func (s *stack) size() int {
	if s.next == nil {
		return 0
	}
	size := 0
	for node := s.next; node != nil; node = node.next {
		size++
	}
	return size
}

func (s *stack) clear() {
	s.next = nil
}

func (s *stack) Out_all_El() {

	if s.next == nil {
		panic("Stack is empty")
	}

	fmt.Print("[")
	for node := s.next; node != nil; node = node.next {
		if node.next == nil {
			fmt.Print(node.data)
			break
		}
		fmt.Print(node.data, ", ")
	}
	fmt.Print("]")
}
