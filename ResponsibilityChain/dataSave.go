package main

import "fmt"

type DataSaveService struct {
	Name string
	Next Service
}

func (dataSaveService *DataSaveService) Execute(d *Data) {
	if !d.UpdateSource {
		fmt.Print("Данные не были обновлены\n")
		dataSaveService.Next.Execute(d)
		return
	}
	fmt.Print("Данные были сохранены\n")

}

func (dataSaveService *DataSaveService) SetNext(s Service) {
	dataSaveService.Next = s
}
