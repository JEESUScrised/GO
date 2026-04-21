package main

import (
	"fmt"
	"sync"
)

type Task struct {
	Number int
}

type Result struct {
	Number    int
	Factorial uint64
}

func factorial(n int) uint64 {
	result := uint64(1)
	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}
	return result
}

func worker(tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		value := factorial(task.Number)
		results <- Result{
			Number:    task.Number,
			Factorial: value,
		}
	}
}

func main() {
	taskCh := make(chan Task, 20)
	resultCh := make(chan Result, 20)

	var workerWG sync.WaitGroup
	workerCount := 4

	for i := 0; i < workerCount; i++ {
		workerWG.Add(1)
		go worker(taskCh, resultCh, &workerWG)
	}

	for i := 1; i <= 20; i++ {
		taskCh <- Task{Number: i}
	}
	close(taskCh)

	workerWG.Wait()
	close(resultCh)

	ordered := make([]uint64, 21)
	for res := range resultCh {
		ordered[res.Number] = res.Factorial
	}

	for i := 1; i <= 20; i++ {
		fmt.Printf("Factorial of %d is %d\n", i, ordered[i])
	}
}
