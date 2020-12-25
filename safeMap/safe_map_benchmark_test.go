package safeMap

import (
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"testing"
)

func BenchmarkSafeMap(b *testing.B) {
	maps := map[string]SafeMap{
		"ShardMap1":  NewSafeShardStr2Any(1),
		"ShardMap16": NewSafeShardStr2Any(16),
		"ShardMap32": NewSafeShardStr2Any(32),
		"ShardMap64": NewSafeShardStr2Any(64),
		"MutexMap":   NewMutexStr2Any(),
	}

	size := 1024

	for _, cpus := range []int{4, 8, 32, 1024} {
		runtime.GOMAXPROCS(cpus)
		for name, m := range maps {
			b.Run(name+"#"+strconv.Itoa(cpus), func(b *testing.B) {
				b.ResetTimer()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						v := rand.Intn(size)
						k := strconv.Itoa(v)
						m.Set(k, v)
						value, ok := m.Get(k)
						if value.(int) != v || !ok {
							log.Fatalf("m.Get value is %v", value)
						}
						/*				switch v % 4 {
										case 0:
											m.Set(k, v)
										case 1:
											m.Get(k)
										case 2:
											m.Delete(k)
										case 3:
											m.Pop(k)
										}*/
					}
				})
			})
		}
	}

}
