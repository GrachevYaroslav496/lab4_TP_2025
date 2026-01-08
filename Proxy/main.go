package main

import (
	"OtherCode/TP/lab4_TP_2025/Proxy"
	"fmt"
)

func main() {
	proxy := Proxy.NewProxy()
	result := proxy.Request()
	fmt.Printf("Результат: %s\n", result)
}

