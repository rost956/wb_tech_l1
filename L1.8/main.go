package main

// Программа устанавливает i-й бит введенного числа в указанное состояние
import (
	"fmt"
)

func setBit(x int64, i uint, bit int) int64 {
	mask := int64(1) << i
	if bit == 1 {
		return x | mask
	}
	return x &^ mask
}

func main() {
	var x int64
	var iInput int
	var b int

	fmt.Print("Введите число (int64): ")
	if _, err := fmt.Scan(&x); err != nil {
		fmt.Println("Ошибка ввода числа:", err)
		return
	}

	fmt.Print("Введите индекс бита i (0..63): ")
	if _, err := fmt.Scan(&iInput); err != nil {
		fmt.Println("Ошибка ввода индекса:", err)
		return
	}
	if iInput < 0 || iInput > 63 {
		fmt.Println("i должно быть от 0 до 63")
		return
	}

	fmt.Print("Введите значение бита (0 или 1): ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Ошибка ввода значения бита:", err)
		return
	}
	if b != 0 && b != 1 {
		fmt.Println("Значение бита должно быть 0 или 1")
		return
	}

	before := x
	after := setBit(x, uint(iInput), b)

	fmt.Printf("Было:  %d (%b)\n", before, before)
	fmt.Printf("Стало: %d (%b)\n", after, after)
}
