package contexttest

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestCounter(t *testing.T) {
	var counter int32     // 计数值
	var wg sync.WaitGroup // 等待 goroutine 执行完毕
	var mu sync.Mutex     // 互斥锁保护共享资源

	for i := 0; i < 100; i++ {
		wg.Go(func() {
			for i := 0; i < 100; i++ {
				mu.Lock()
				counter++ // 加锁保护，避免竞争条件
				mu.Unlock()
			}
		})
	}
	wg.Wait()
	fmt.Println(counter)
}

func TestCounterAtomic(t *testing.T) {
	var counter int64     // 原子操作计数值
	var wg sync.WaitGroup // 等待 goroutine 执行完毕

	for i := 0; i < 100; i++ {
		wg.Go(func() {
			for i := 0; i < 100; i++ {
				atomic.AddInt64(&counter, 1) // 原子操作，无需加锁
			}
		})
	}
	wg.Wait()
	result := atomic.LoadInt64(&counter)
	fmt.Printf("Atomic counter result: %d\n", result)
}

// 基准测试：互斥锁方式
func BenchmarkCounterMutex(b *testing.B) {
	var counter int32
	var wg sync.WaitGroup
	var mu sync.Mutex

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		wg.Go(func() {
			mu.Lock()
			counter++
			mu.Unlock()
		})
	}
	wg.Wait()
}

// 基准测试：原子操作方式
func BenchmarkCounterAtomic(b *testing.B) {
	var counter int64
	var wg sync.WaitGroup

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		wg.Go(func() {
			atomic.AddInt64(&counter, 1)
		})
	}
	wg.Wait()
}
