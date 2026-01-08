package Adapter

import "fmt"

type Adaptee struct {
}

func NewAdaptee() *Adaptee {
	return &Adaptee{}
}

func (a *Adaptee) SpecificRequest() string {
	fmt.Println("Adaptee: выполнение специфического запроса")
	return "Результат от Adaptee"
}

