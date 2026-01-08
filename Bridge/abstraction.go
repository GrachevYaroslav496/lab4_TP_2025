package Bridge

import "fmt"

type Abstraction struct {
	implementer Implementer
}

func NewAbstraction(impl Implementer) *Abstraction {
	return &Abstraction{
		implementer: impl,
	}
}

func (a *Abstraction) Operation() string {
	fmt.Println("Абстракция: выполнение операции")
	return a.implementer.OperationImpl()
}

type RefinedAbstraction struct {
	*Abstraction
}

func NewRefinedAbstraction(impl Implementer) *RefinedAbstraction {
	return &RefinedAbstraction{
		Abstraction: NewAbstraction(impl),
	}
}

func (ra *RefinedAbstraction) Operation() string {
	fmt.Println("Уточненная абстракция: дополнительная обработка")
	return ra.Abstraction.Operation()
}

