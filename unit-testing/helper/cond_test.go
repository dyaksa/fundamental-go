package helper

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("done : ", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for j := 0; j < 10; j++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// go func() {
	// 	for j := 0; j < 10; j++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Broadcast()
	// 	}
	// }()

	group.Wait()
}
