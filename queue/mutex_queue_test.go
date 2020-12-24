package queue

import (
	"testing"
)

func TestMQueue_EnqueueAndDequeue(t *testing.T) {
	mQueue := NewMQueue(10)
	length := 100
	vs := make([]int, 0, length)
	for i := 0; i < length; i++ {
		vs = append(vs, i)
	}

	for _, value := range vs {
		mQueue.Enqueue(value)
	}

	for _, value := range vs {
		v := mQueue.Dequeue()
		newV := v.(int)
		if newV != value {
			t.Fail()
			t.Logf("the enqueue value is %d but dequeue value is %d", value, newV)
		}
	}

}
