package bench_test

import (
	"math/rand"
	"testing"

	"github.com/jonhoo/drwmutex"
)

func BenchmarkIncrement(b *testing.B) {
	var v int
	for i := 0; i < b.N; i++ {
		v = v + 1
	}
}

func BenchmarkRandFloat64Instance(b *testing.B) {
	random := rand.New(rand.NewSource(rand.Int63()))
	for i := 0; i < b.N; i++ {
		_ = random.Float64()
	}
}

func BenchmarkRandFloat64InstanceClear(b *testing.B) {
	random := rand.New(rand.NewSource(rand.Int63()))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = random.Float64()
	}
}

func BenchmarkCPUID(b *testing.B) {
	mu := drwmutex.New()
	for i := 0; i < b.N; i++ {
		_ = mu.RLocker()
	}
}

func BenchmarkCPUIDClear(b *testing.B) {
	mu := drwmutex.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mu.RLocker()
	}
}

func BenchmarkLockUnlock(b *testing.B) {
	mu := drwmutex.New()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}

func BenchmarkLockUnlockClear(b *testing.B) {
	mu := drwmutex.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}

func BenchmarkG(b *testing.B) {
	b.SetParallelism(1)
	b.RunParallel(func(p *testing.PB) {
		var v int
		for p.Next() {
			for i := 0; i < 1000; i++ {
				v = v + 1
			}
		}
	})
}

func BenchmarkMap10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = map[string]string{
			"0": "9",
			"1": "8",
			"2": "7",
			"3": "6",
			"4": "5",
			"5": "4",
			"6": "3",
			"7": "2",
			"8": "1",
			"9": "0",
		}
	}
}

func BenchmarkS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = news()
	}
}

func news() string {
	return "kkk"
}

func BenchmarkMap10Raw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[string]string, 10)
	}
}
