package safeQueue

type Queue interface {
	Enqueue(v interface{})
	Dequeue() interface{}
}
