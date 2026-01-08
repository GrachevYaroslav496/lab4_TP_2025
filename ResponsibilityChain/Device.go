package main

import "fmt"

type Device struct {
	Name string
	Next Service
}

func (device *Device) Execute(d *Data) {
	if d.GetSource {
		fmt.Print("Данные были получены ранее\n")
		device.Next.Execute(d)
		return
	}
	fmt.Printf("Получены данные из %s\n", device.Name)
	d.GetSource = true
	device.Next.Execute(d)

}

func (device *Device) SetNext(s Service) {
	device.Next = s
}
