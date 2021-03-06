package test

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextD.Value("b"))
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))

}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	counter := 1

	go func() {
		defer close(destination)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestWithCancel(t *testing.T) {
	fmt.Println("Total goroutines:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("counter ", n)
		if n == 10 {
			break
		}
	}
	cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("Total goroutines:", runtime.NumGoroutine())
}

func TestWithTimeout(t *testing.T) {
	fmt.Println("Total goroutines:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)

	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("counter ", n)
	}

	time.Sleep(7 * time.Second)
	fmt.Println("Total goroutines:", runtime.NumGoroutine())
}

func TestWithDeadLine(t *testing.T) {
	fmt.Println("Total goroutines:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))

	defer cancel()

	dest := CreateCounter(ctx)

	for n := range dest {
		fmt.Println("counter ", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total goroutines:", runtime.NumGoroutine())
}
