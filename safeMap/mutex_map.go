package safeMap

import "sync"

type MutexStr2Any struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewMutexStr2Any(size ...int) *MutexStr2Any {

	def := _DefaultSize
	if len(size) > 0 && size[0] > 0 {
		def = size[0]
	}

	m := MutexStr2Any{
		data: make(map[string]interface{}, def),
	}

	return &m
}

// Sets the given value under the specified key.
func (m *MutexStr2Any) Set(key string, value interface{}) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()
}

// Get retrieves an element from map under given key.
func (m *MutexStr2Any) Get(key string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	val, ok := m.data[key]
	return val, ok
}

// Delete delete an element from the map.
func (m *MutexStr2Any) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.data, key)
}

//Pop pop the value of by the key. if the key not exits, return false
func (m *MutexStr2Any) Pop(key string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if val, ok := m.data[key]; ok {
		delete(m.data, key)
		return val, true
	}

	return nil, false
}
