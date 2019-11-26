package main

import (
	"fmt"

	"github.com/whuhyw/data_structure/queue"
)

func main() {
	s := queue.NewQueue()
	s.Push("a")
	s.Push(1)
	fmt.Println(s.Pop())
}
