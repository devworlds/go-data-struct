package main

import "fmt"

type Stack struct {
	Tail *Node
}

type Node struct {
	Val  string
	Next *Node
}

func (s *Stack) Push(val string) {
	node := Node{Val: val} // um, dois, tres
	if s.Tail != nil {
		node.Next = s.Tail // dois.next = um, tres.next = dois
	}

	s.Tail = &node //Tail = um, Tail = dois, Tail = 3
}

func (s *Stack) Pop() string {
	if s.Tail == nil {
		return "data not found"
	}

	node := s.Tail
	s.Tail = node.Next
	return node.Val
}

func main() {
	stack := Stack{}
	stack.Push("um")
	stack.Push("dois")
	stack.Push("tres")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}