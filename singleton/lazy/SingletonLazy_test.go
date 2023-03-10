package lazy

import (
	"sync"
	"testing"
)

const parCount = 100

func TestSingletonLazy(t *testing.T) {
	s1 := GetSingelton()
	s2 := GetSingelton()
	s3 := GetSingelton()

	t.Logf("s1 's addr :%p\n", s1)
	t.Logf("s2 's addr :%p\n", s2)
	t.Logf("s3 's addr :%p\n", s3)
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*singelton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			// 协程阻塞，等待channel被关闭才能继续运行
			<-start
			instances[index] = GetSingelton()
			wg.Done()
		}(i)
	}
	// 关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
