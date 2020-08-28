package stack

import "sync"

var _ Stack = (*threadSafeStack)(nil)

type threadSafeStack struct {
	s Stack
	sync.RWMutex
}

func newThreadSafeStack() *threadSafeStack {
	return &threadSafeStack{
		s: NewStack(),
	}
}

func (t *threadSafeStack) Push(e Element) {
	t.RLock()
	t.Lock()
	t.s.Push(e)
	t.RUnlock()
	t.Unlock()
}

func (t *threadSafeStack) Pop() Element {
	t.Lock()
	t.s.Push(e)
	t.Unlock()
}

func (t *threadSafeStack) Top() Element {
	t.RLock()
	t.s.Push(e)
	t.Unlock()
}

func (t *threadSafeStack) Clear() {
	panic("implement me")
}

func (t *threadSafeStack) Size() int {
	panic("implement me")
}

func (t *threadSafeStack) Empty() bool {
	panic("implement me")
}

func (t *threadSafeStack) String() string {
	panic("implement me")
}

func (t *threadSafeStack) Equal(s Stack) bool {
	panic("implement me")
}
