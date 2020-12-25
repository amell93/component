package safeMap

import "sync"

const _ShardCount = 32

// fnv32 32 bit fnv hash algorithm
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// A "thread" safe map of type string:Anything.
// To avoid lock bottlenecks this map is dived to several (__SHARD_COUNT) map shards.
type SafeShardStr2Any struct {
	data      []*safeStrAnyShared
	shardNums int
}

// A "thread" safe string to anything map.
type safeStrAnyShared struct {
	items        map[string]interface{}
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

// NewSafeShardStr2Any create a new concurrent map.
func NewSafeShardStr2Any(shardNums ...int) *SafeShardStr2Any {
	sn := _ShardCount

	if len(shardNums) > 0 && shardNums[0] > 0 {
		sn = shardNums[0]
	}

	cm := &SafeShardStr2Any{
		data:      nil,
		shardNums: sn,
	}

	data := make([]*safeStrAnyShared, cm.shardNums)
	for i := 0; i < cm.shardNums; i++ {
		data[i] = &safeStrAnyShared{items: make(map[string]interface{}, _DefaultSize)}
	}
	cm.data = data
	return cm
}

// getShard returns shard under given key
func (m *SafeShardStr2Any) getShard(key string) *safeStrAnyShared {
	return m.data[uint(fnv32(key))%uint(m.shardNums)]
}

// Sets the given value under the specified key.
func (m *SafeShardStr2Any) Set(key string, value interface{}) {
	// Get map shard.
	shard := m.getShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

// Get retrieves an element from map under given key.
func (m *SafeShardStr2Any) Get(key string) (interface{}, bool) {
	// Get shard
	shard := m.getShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// Delete delete an element from the map.
func (m *SafeShardStr2Any) Delete(key string) {
	// Try to get shard.
	shard := m.getShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

//Pop pop the value of by the key. if the key not exits, return false
func (m *SafeShardStr2Any) Pop(key string) (interface{}, bool) {

	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()

	if val, ok := shard.items[key]; ok {
		delete(shard.items, key)
		return val, true
	}

	return nil, false
}
