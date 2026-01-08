package main

func main() {
	device := &Device{Name: "Device"}
	updateService := &UpdateDataService{Name: "UpdateDataService"}
	dataService := &DataSaveService{Name: "DataSaveService"}

	device.SetNext(updateService)
	updateService.SetNext(dataService)

	data := &Data{}
	device.Execute(data)
}
