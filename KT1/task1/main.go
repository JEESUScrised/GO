package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

// печать результата: целые без дробной части, остальное — с нужной точностью
func ratString(r *big.Rat) string {
	if r.Sign() == 0 {
		return "0"
	}
	if r.IsInt() {
		return new(big.Int).Set(r.Num()).String()
	}
	var f big.Float
	f.SetPrec(4096)
	f.SetRat(r)
	s := f.Text('f', -1)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	if s == "" || s == "-" {
		return "0"
	}
	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)
	readLine := func() (string, error) {
		line, err := in.ReadString('\n')
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(line), nil
	}

	var a, b big.Rat

	fmt.Print("Введите первое число: ")
	s1, err := readLine()
	if err != nil || s1 == "" {
		fmt.Println("Ошибка: пустой ввод")
		return
	}
	if _, ok := a.SetString(s1); !ok {
		fmt.Println("Ошибка: это не число")
		return
	}

	fmt.Print("Введите второе число: ")
	s2, err := readLine()
	if err != nil || s2 == "" {
		fmt.Println("Ошибка: пустой ввод")
		return
	}
	if _, ok := b.SetString(s2); !ok {
		fmt.Println("Ошибка: это не число")
		return
	}

	fmt.Print("Операция (+, -, *, /): ")
	op, err := readLine()
	if err != nil || op == "" {
		fmt.Println("Ошибка: операция не введена")
		return
	}

	var res big.Rat
	switch op {
	case "+":
		res.Add(&a, &b)
	case "-":
		res.Sub(&a, &b)
	case "*":
		res.Mul(&a, &b)
	case "/":
		if b.Sign() == 0 {
			fmt.Println("Ошибка: деление на ноль нельзя")
			return
		}
		res.Quo(&a, &b)
	default:
		fmt.Println("Неизвестная операция, используйте + - * /")
		return
	}

	fmt.Println("Результат:", ratString(&res))
}
