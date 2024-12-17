package main

type Queue struct {
	Head *Node
}

type Node struct {
	Val  string
	Next *Node
}

func (q *Queue) Enqueue(val string) {
	node := &Node{Val: val}
	if q.Head == nil {
		q.Head = node
	} else {
		curr := q.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
	}
}

func (q *Queue) Dequeue() string {
	if q.Head == nil {
		return "no queue messages"
	}

	node := q.Head
	q.Head = q.Head.Next

	return node.Val
}
