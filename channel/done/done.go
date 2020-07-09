package main

import (
	"fmt"
	"sync"
)

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func doWork(id int, w worker) {
	func() {
		for n := range w.in {
			fmt.Printf("Worker %d received %c\n", id, n)
			w.done()
		}
	}()
}

type worker struct {
	in   chan int
	done func()
}

func chanDemo() {
	var wg sync.WaitGroup
	//wg.Add(20)

	var workers [10]worker
	for i, _ := range workers {
		workers[i] = createWorker(i, &wg)
	}

	for i, v := range workers {
		v.in <- 'a' + i
		wg.Add(1)
	}

	for i, v := range workers {
		v.in <- 'A' + i
		wg.Add(1)
	}

	wg.Wait()

	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
}

func main() {
	fmt.Println("Channel close")
	chanDemo()
}
