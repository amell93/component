package unsafe

import "unsafe"

// src/reflect/value.go
//type StringHeader struct {
//	Data uintptr
//	Len int
//}

// src/reflect/value.go
//type SliceHeader struct {
//	Data uintptr
//	Len int
//	Cap int
//}

// String2Slice convert string to []byte by zero-copy
func String2Slice(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func Slice2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
