package Bridge

import "fmt"

type ConcreteImplementerB struct {
}

func NewConcreteImplementerB() *ConcreteImplementerB {
	return &ConcreteImplementerB{}
}

func (cib *ConcreteImplementerB) OperationImpl() string {
	fmt.Println("Реализация B: выполнение операции")
	return "Результат реализации B"
}

