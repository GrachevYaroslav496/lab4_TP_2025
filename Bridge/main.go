package main

import (
	"OtherCode/TP/lab4_TP_2025/Bridge"
	"fmt"
)

func main() {
	implA := Bridge.NewConcreteImplementerA()
	abstractionA := Bridge.NewAbstraction(implA)
	resultA := abstractionA.Operation()
	fmt.Printf("Результат A: %s\n\n", resultA)

	implB := Bridge.NewConcreteImplementerB()
	refinedAbstraction := Bridge.NewRefinedAbstraction(implB)
	resultB := refinedAbstraction.Operation()
	fmt.Printf("Результат B: %s\n", resultB)
}

