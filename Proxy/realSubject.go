package Proxy

import "fmt"

type RealSubject struct {
}

func NewRealSubject() *RealSubject {
	return &RealSubject{}
}

func (rs *RealSubject) Request() string {
	fmt.Println("Выполнение запроса в RealSubject")
	return "Результат от RealSubject"
}

