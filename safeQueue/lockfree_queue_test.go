package safeQueue

import "testing"

func TestLFQueue_EnqueueAndDequeue(t *testing.T) {
	lfQueue := NewLFQueue()
	length := 100
	vs := make([]int, 0, length)
	for i := 0; i < length; i++ {
		vs = append(vs, i)
	}

	for _, value := range vs {
		lfQueue.Enqueue(value)
	}

	for _, value := range vs {
		v := lfQueue.Dequeue()
		newV := v.(int)
		if newV != value {
			t.Fail()
			t.Logf("the enqueue value is %d but dequeue value is %d", value, newV)
		}
	}
}
