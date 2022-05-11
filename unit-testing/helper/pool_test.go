package helper

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	group := sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}

	pool.Put("Dyaksa")
	pool.Put("Jauharuddin")
	pool.Put("Nour")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
		group.Done()
	}

	group.Wait()
	fmt.Println("Selesai")
}
