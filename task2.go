package main

import "fmt"

func generator(out chan<- int) {
	defer close(out)

	for i := 1; i <= 10; i++ {
		out <- i
	}
}

func doubler(in <-chan int, out chan<- int) {
	defer close(out)

	for n := range in {
		out <- n * 2
	}
}

func printer(in <-chan int, done chan<- struct{}) {
	defer close(done)

	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	genCh := make(chan int, 10)
	doubleCh := make(chan int, 10)
	done := make(chan struct{})

	go generator(genCh)
	go doubler(genCh, doubleCh)
	go printer(doubleCh, done)

	<-done
}
