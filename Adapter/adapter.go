package Adapter

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter() *Adapter {
	return &Adapter{
		adaptee: NewAdaptee(),
	}
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

