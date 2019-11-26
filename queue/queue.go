package queue

type Element interface{}

type Queue interface {
	Push(e Element)
	Pop() Element
	Peek() Element
	Size() int
	Empty() bool
	Clear()
	String() string
}