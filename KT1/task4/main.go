package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var arr [10]int
	for i := range arr {
		arr[i] = rand.Intn(100) + 1
	}

	fmt.Print("Массив: ")
	for i, v := range arr {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println()

	sum := 0
	minVal := arr[0]
	maxVal := arr[0]
	for _, v := range arr {
		sum += v
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	avg := float64(sum) / float64(len(arr))

	fmt.Println("Сумма:", sum)
	fmt.Println("Минимум:", minVal)
	fmt.Println("Максимум:", maxVal)
	fmt.Printf("Среднее: %.2f\n", avg)
}
