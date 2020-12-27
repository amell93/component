package unsafe

import (
	"log"
	"testing"
)

var (
	s1  = "abcdefghijklmnopqrstuvwxyz"
	bs1 = []byte(s1)

	s2  = "abcdefghijkl中国mnopqrstuvwxyz"
	bs2 = []byte(s2)
)

func TestSlice2String(t *testing.T) {
	newS1 := Slice2String(bs1)
	if s1 != newS1 {
		log.Fatalf("Slice2String err")
	}

	newS2 := Slice2String(bs2)
	if s2 != newS2 {
		log.Fatalf("Slice2String err")
	}

}

func TestString2Slice(t *testing.T) {
	newBs1 := String2Slice(s1)
	for index, b := range newBs1 {
		if b != bs1[index] {
			log.Fatalf("String2Slice failed")
		}
	}

	newBs2 := String2Slice(s2)
	for index, b := range newBs2 {
		if b != bs2[index] {
			log.Fatalf("String2Slice failed")
		}
	}

}

func BenchmarkSlice2String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := Slice2String(bs1)
		_ = s
	}
}

func BenchmarkString2Slice(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs := String2Slice(s)
		_ = bs
	}
}
