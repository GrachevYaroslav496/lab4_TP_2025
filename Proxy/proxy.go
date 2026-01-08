package Proxy

import "fmt"

type Proxy struct {
	realSubject *RealSubject
}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (p *Proxy) Request() string {
	if p.realSubject == nil {
		fmt.Println("Создание RealSubject через Proxy")
		p.realSubject = NewRealSubject()
	}
	fmt.Println("Proxy обрабатывает запрос перед передачей")
	result := p.realSubject.Request()
	fmt.Println("Proxy обрабатывает результат после получения")
	return result
}

