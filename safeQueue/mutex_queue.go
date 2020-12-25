package safeQueue

import (
	"sync"
)

type MQueue struct {
	data []interface{}
	mu   sync.Mutex
}

const defaultSize = 32

func NewMQueue(size ...int) *MQueue {

	length := defaultSize

	if len(size) > 0 && size[0] >= 0 {
		length = size[0]
	}

	mq := MQueue{
		data: make([]interface{}, 0, length),
	}

	return &mq
}

func (q *MQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

func (q *MQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}

	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}
