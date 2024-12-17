package main

import "fmt"

func main() {
	queue := Queue{}

	queue.Enqueue("message1")
	queue.Enqueue("message2")
	queue.Enqueue("message3")

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
