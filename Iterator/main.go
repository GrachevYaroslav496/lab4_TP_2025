package main

import (
	"OtherCode/TP/lab4_TP_2025/Iterator"
	"fmt"
)

func main() {
	fmt.Println("=== Итератор для массива целых чисел ===")
	numbers := []int{1, 2, 3, 4, 5}
	intCollection := Iterator.NewIntArrayCollection(numbers)
	intIterator := intCollection.CreateIterator()

	for intIterator.HasNext() {
		value := intIterator.Next()
		fmt.Printf("Число: %d\n", value)
	}

	fmt.Println("\n=== Итератор для слайса строк ===")
	names := []string{"Алиса", "Боб", "Чарли", "Диана"}
	stringCollection := Iterator.NewStringSliceCollection(names)
	stringIterator := stringCollection.CreateIterator()

	for stringIterator.HasNext() {
		value := stringIterator.Next()
		fmt.Printf("Имя: %s\n", value)
	}
}
