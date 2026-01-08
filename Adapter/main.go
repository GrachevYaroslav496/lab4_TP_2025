package main

import (
	"OtherCode/TP/lab4_TP_2025/Adapter"
	"fmt"
)

func main() {
	adapter := Adapter.NewAdapter()
	result := adapter.Request()
	fmt.Printf("Результат через адаптер: %s\n", result)
}

