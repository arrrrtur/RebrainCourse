package test

import (
	"module5/07_task/cmd"
	"sync"
	"testing"
)

func BenchmarkWithoutPool(b *testing.B) {
	firstCase := cmd.Counter{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10_000; j++ {
			b.StopTimer()
			firstCase.Inc()
			b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var counterPool = sync.Pool{New: func() interface{} { return new(cmd.Counter) }}
	secondCase := counterPool.Get().(*cmd.Counter)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10_000; j++ {
			b.StopTimer()
			secondCase.Inc()
			b.StartTimer()
		}
	}
}
