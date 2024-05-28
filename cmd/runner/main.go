package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/jonhoo/drwmutex"
)

type Lockable interface {
	sync.Locker
	RLocker() sync.Locker
}

func main() {
	_g := flag.Int("g", 0, "количество горутин на cpu")
	_p := flag.Float64("p", 0, "вероятность записи в каждой итерации")
	_P := flag.Float64("P", 0, "вероятность вызова CPUID в каждой итерации на чтение")
	_r := flag.Int("r", 0, "нагрузка при чтении")
	_w := flag.Int("w", 0, "нагрузка при записи")
	_x := flag.Int("x", 0, "использовать распределенный мьютекс")

	flag.Parse()

	g := *_g
	p := *_p
	P := *_P
	r := *_r
	w := *_w
	x := *_x

	m := runtime.GOMAXPROCS(0)

	P_over_p := p + (1-p)*P

	i_all := g * m * 100000
	if p != 1 && P != 0 {
		if another := float64((m * 100)) / ((1 - p) * P); float64(i_all) < another {
			i_all = int(math.Ceil(another))
		}
	}
	if p != 0 {
		if another := float64((m * 100)) / p; float64(i_all) < another {
			i_all = int(math.Ceil(another))
		}
	}
	i_g := i_all / (g * m)

	mg_range := make([]struct{}, m*g)
	r_range := make([]struct{}, r)
	w_range := make([]struct{}, w)

	var mu Lockable
	switch x {
	case 0:
		mu = &sync.RWMutex{}
	case 1:
		mu = drwmutex.New()
	}

	wg := &sync.WaitGroup{}
	wg.Add(g * m)

	start := time.Now()

	for range mg_range {
		go func() {
			defer wg.Done()

			var v int64
			var e float64

			readLocker := mu.RLocker()
			random := rand.New(rand.NewSource(rand.Int63()))

			for i := 0; i < i_g; i++ {
				e = random.Float64()
				if e < p {
					mu.Lock()
					for range w_range {
						v = v + 1
					}
					mu.Unlock()
					continue
				}
				if e < P_over_p {
					readLocker = mu.RLocker()
				}
				readLocker.Lock()
				for range r_range {
					v = v + 1
				}
				readLocker.Unlock()
			}
		}()
	}

	wg.Wait()
	end := time.Now()

	var l int
	if drw, ok := mu.(drwmutex.DRWMutex); ok {
		l = len(drw)
	}

	fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n", g, i_all, p, P, r, w, runtime.NumCPU(), m, x, l, end.Sub(start).Seconds())
}
