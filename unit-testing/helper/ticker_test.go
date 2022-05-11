package helper

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	timer := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		timer.Stop()
	}()

	for time := range timer.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}
