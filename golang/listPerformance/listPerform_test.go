package list_test

import (
	"container/list"
	"testing"
	"time"
)

func BenchmarkNoTick(b *testing.B) {
	nums := list.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nums.PushBack(i)
	}
}

func BenchmarkTickMS(b *testing.B) {
	nums := list.New()
	b.ResetTimer()
	c := time.Tick(1 * time.Millisecond)
	for i := 0; i < b.N; i++ {
		_ = <-c
		nums.PushBack(i)
	}
}

func BenchmarkTickUS(b *testing.B) {
	nums := list.New()
	b.ResetTimer()
	c := time.Tick(1 * time.Microsecond)
	for i := 0; i < b.N; i++ {
		_ = <-c
		nums.PushBack(i)
	}
}

func BenchmarkTickNS(b *testing.B) {
	nums := list.New()
	b.ResetTimer()
	c := time.Tick(1 * time.Nanosecond)
	for i := 0; i < b.N; i++ {
		_ = <-c
		nums.PushBack(i)
	}
}
