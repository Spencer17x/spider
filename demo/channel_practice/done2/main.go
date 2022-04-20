package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

var count int

func doWorker(id int, ch <-chan int, done chan<- bool) {
	for n := range ch {
		count++
		fmt.Printf("Worker %d received %c\n", id, n)
		done <- true
	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		<-workers[i].done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}

	//time.Sleep(time.Millisecond)

	fmt.Println("count:", count)
}

func main() {
	chanDemo()
}
