package queue

import (
	"github.com/whuhyw/data_structure/stack"
)

var _ Queue = (*entry)(nil)

func NewQueue() *entry {
	return newEntry()
}

func newEntry() *entry {
	return &entry{
		s1:   stack.NewStack(),
		s2:   stack.NewStack(),
		size: 0,
	}
}

type entry struct {
	s1   stack.Stack
	s2   stack.Stack
	size int
}

func (e *entry) Clear() {
	*e = *newEntry()
}

func (e *entry) reload() {
	if e.Empty() || e.s2.Size() > 0 {
		return
	}
	for e.s1.Size() > 0 {
		e.s2.Push(e.s1.Pop())
	}
}

func (e *entry) Push(element Element) {
	e.size++
	e.s1.Push(element)
}

func (e *entry) Pop() Element {
	if e.Empty() {
		return nil
	}
	if e.s2.Size() == 0 {
		e.reload()
	}
	e.size--
	return e.s2.Pop()
}

func (e *entry) Peek() Element {
	if e.Empty() {
		return nil
	}
	e.reload()
	return e.s2.Top()
}

func (e *entry) Size() int {
	return e.size
}

func (e *entry) Empty() bool {
	return e.size == 0
}

func (e *entry) String() string {
	s := NewQueue()
	for !e.Empty() {
		s.Push(e.Pop())
	}
	return s.String()
}
