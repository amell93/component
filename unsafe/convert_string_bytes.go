package unsafe

import "unsafe"

// src/reflect/value.go
//type StringHeader struct {
//	Data uintptr
//	Len int
//}

// src/runtime/string.go
//type stringStruct struct {
//	str unsafe.Pointer
//	len int
//}

// src/reflect/value.go
//type SliceHeader struct {
//	Data uintptr
//	Len int
//	Cap int
//}

// src/runtime/slice.go
//type slice struct {
//	array unsafe.Pointer
//	len   int
//	cap   int
//}

// String2Slice convert string to []byte by zero-copy.
// waring:
//     program will panic if when modify the return []byte.
func String2Slice(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// String2Slice convert []byte to string by zero-copy.
// waring:
//     the s string will be modified when you modify the data of the []byte b.
func Slice2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
