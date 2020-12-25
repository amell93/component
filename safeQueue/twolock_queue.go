package safeQueue

import "sync"

// TLQueue is a concurrent unbounded queue.
type TLQueue struct {
	head  *tlNode
	tail  *tlNode
	hLock sync.Mutex //protect head
	tLock sync.Mutex //protect tail

}

type tlNode struct {
	value interface{}
	next  *tlNode
}

// NewTLQueue returns an empty TLQueue.
func NewTLQueue() *TLQueue {
	n := &tlNode{}
	tq := TLQueue{head: n, tail: n}

	return &tq
}

// Enqueue puts the given value v at the tail of the queue.
func (q *TLQueue) Enqueue(v interface{}) {
	n := &tlNode{value: v}
	q.tLock.Lock()
	q.tail.next = n
	q.tail = n
	q.tLock.Unlock()
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *TLQueue) Dequeue() interface{} {
	q.hLock.Lock()
	n := q.head

	newHead := n.next
	if newHead == nil {
		q.hLock.Unlock()
		return nil
	}

	v := newHead.value
	newHead.value = nil
	q.head = newHead
	q.hLock.Unlock()

	return v
}
