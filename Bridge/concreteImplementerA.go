package Bridge

import "fmt"

type ConcreteImplementerA struct {
}

func NewConcreteImplementerA() *ConcreteImplementerA {
	return &ConcreteImplementerA{}
}

func (cia *ConcreteImplementerA) OperationImpl() string {
	fmt.Println("Реализация A: выполнение операции")
	return "Результат реализации A"
}

