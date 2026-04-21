package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1

	fmt.Println("Я загадал число от 1 до 100. Угадайте!")
	var guess int

	for {
		fmt.Print("Ваш вариант: ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Введите целое число")
			continue
		}

		if guess < 1 || guess > 100 {
			fmt.Println("Число должно быть в диапазоне 1..100")
			continue
		}

		if guess == secret {
			fmt.Println("Поздравляю, вы угадали!")
			break
		}
		if guess < secret {
			fmt.Println("Больше")
		} else {
			fmt.Println("Меньше")
		}
	}
}
