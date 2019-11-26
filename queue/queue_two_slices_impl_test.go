package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_entry(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 0, q.Size())
}
