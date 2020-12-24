package queue

import (
	"sync/atomic"
	"unsafe"
)

// LFQueue is a lock-free unbounded queue.
type LFQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

type lfNode struct {
	value interface{}
	next  unsafe.Pointer
}

// NewLFQueue returns an empty LFQueue.
func NewLFQueue() *LFQueue {
	n := unsafe.Pointer(&lfNode{})

	return &LFQueue{head: n, tail: n}
}

func (q *LFQueue) Enqueue(v interface{}) {
	n := &lfNode{value: v}
	for {
		tailNode := load(&q.tail)        //tail snapshot
		nextNode := load(&tailNode.next) //tail.nextNode snapshot
		if tailNode == load(&q.tail) {   //check consistent between tail snapshot and current tail
			if nextNode == nil { //check tail.nextNode
				if cas(&tailNode.next, nextNode, n) { //update the tail.nextNode
					cas(&q.tail, tailNode, n) //update the tail
					return
				}
			} else { // tail.nextNode has been update
				cas(&q.tail, tailNode, nextNode) //try to swing tail to the nextNode node
			}
		}
	}
}

func (q *LFQueue) Dequeue() interface{} {
	for {
		headNode := load(&q.head)
		tailNode := load(&q.tail)
		nextNode := load(&headNode.next)
		if headNode == load(&q.head) { // are headNode, tailNode, and nextNode consistent?
			if headNode == tailNode { // is queue empty or tailNode falling behind?
				if nextNode == nil { // is queue empty?
					return nil
				}
				// tailNode is falling behind.  try to advance it
				cas(&q.tail, tailNode, nextNode)
			} else {
				// read value before CAS otherwise another dequeue might free the nextNode node
				v := nextNode.value
				if cas(&q.head, headNode, nextNode) {
					return v // Dequeue is done.  return
				}
			}
		}
	}
}

func load(p *unsafe.Pointer) (n *lfNode) {
	return (*lfNode)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *lfNode) (ok bool) {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new))
}
