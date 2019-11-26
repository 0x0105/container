package main

import (
	"fmt"

	"aaron.local/queue"
)

func main() {
	s := queue.NewQueue()
	s.Push("a")
	s.Push(1)
	fmt.Println(s.Pop())
}
