// Package syncMap provides some string->interface{} map containers which is concurrent-safe
package safeMap

const _DefaultSize = 128

type SafeMap interface {
	Get(key string) (value interface{}, ok bool)
	Set(key string, value interface{})
	Delete(key string)
	Pop(key string) (value interface{}, ok bool)
}
