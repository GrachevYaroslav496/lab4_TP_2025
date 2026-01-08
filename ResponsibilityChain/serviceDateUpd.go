package main

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (updateDataService *UpdateDataService) Execute(d *Data) {
	if d.UpdateSource {
		fmt.Print("Данные были получены ранее\n")
		updateDataService.Next.Execute(d)
		return
	}
	fmt.Printf("Получены данные из %s\n", updateDataService.Name)
	d.UpdateSource = true
	updateDataService.Next.Execute(d)

}

func (updateDataService *UpdateDataService) SetNext(s Service) {
	updateDataService.Next = s
}
