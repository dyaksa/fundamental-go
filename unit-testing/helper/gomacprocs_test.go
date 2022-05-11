package helper

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomacProcs(t *testing.T) {

	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Go Routine : ", totalGoRoutine)

	group.Wait()
}
