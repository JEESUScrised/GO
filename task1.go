package main

import (
	"fmt"
	"sync"
)

func partialSum(numbers []int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	out <- sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mid := len(numbers) / 2
	leftPart := numbers[:mid]
	rightPart := numbers[mid:]

	results := make(chan int, 2)
	var wg sync.WaitGroup

	wg.Add(2)
	go partialSum(leftPart, results, &wg)
	go partialSum(rightPart, results, &wg)

	wg.Wait()
	close(results)

	total := 0
	for value := range results {
		total += value
	}

	fmt.Println("Sum:", total)
}
