package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Введите строку: ")
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения")
		return
	}
	s := strings.TrimSpace(line)
	if s == "" {
		fmt.Println("Вы ничего не ввели")
		return
	}

	pal := true
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			pal = false
			break
		}
	}

	if pal {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}
