package helper

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Dyaksa Jauharuddin Nour"
		fmt.Println("selesai mengisi data channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dyaksa Jauharuddin Nour"
	fmt.Println("selesai mengisi data dalam channel")
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Dyaksa Jauharuddin Nour"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestOnlyInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Dyaksa"
		channel <- "Jauharuddin"
		channel <- "Nour"
	}()

	go func() {
		fmt.Println(<-channel)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data " + data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go func() {
		channel1 <- "Dyaksa"
		channel2 <- "Jauharuddin"
	}()

	go func() {
		counter := 0
		for {
			select {
			case data := <-channel1:
				fmt.Println("data ke ", data)
				counter++
			case data := <-channel2:
				fmt.Println("data ke ", data)
				counter++
			}
			if counter == 2 {
				break
			}
		}
	}()

	time.Sleep(2 * time.Second)
}

func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go func() {
		channel1 <- "Dyaksa"
		channel2 <- "Jauharuddin"
	}()

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data ke ", data)
			counter++
		case data := <-channel2:
			fmt.Println("data ke ", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
