package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	func() {
		for n := range c {
			fmt.Printf("Worker %d received %d\n", id, n)
		}
	}()
}

func chanDemo() {
	var channels [10]chan<- int
	for i, _ := range channels {
		channels[i] = createWorker(i)
	}

	for i, v := range channels {
		v <- 'a' + i
	}

	for i, v := range channels {
		v <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Microsecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Microsecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	//channelClose()

	fmt.Println("Buffered channel")
	//bufferedChannel()

	fmt.Println("Channel close")
	chanDemo()
}
