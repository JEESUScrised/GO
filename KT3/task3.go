package main

import (
	"fmt"
	"strings"
	"sync"
)

func countWordsPart(part []string, total *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	localCount := 0
	for _, word := range part {
		if word != "" {
			localCount++
		}
	}

	mu.Lock()
	*total += localCount
	mu.Unlock()
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	words := strings.Split(text, " ")

	workers := 3
	chunkSize := (len(words) + workers - 1) / workers

	var wg sync.WaitGroup
	var mu sync.Mutex
	totalWords := 0

	for i := 0; i < len(words); i += chunkSize {
		end := i + chunkSize
		if end > len(words) {
			end = len(words)
		}

		wg.Add(1)
		go countWordsPart(words[i:end], &totalWords, &mu, &wg)
	}

	wg.Wait()
	fmt.Println("Number of words:", totalWords)
}
