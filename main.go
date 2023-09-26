package main

func main() {
	filePath := "services.json"
	services := ReadServicesFromJson(filePath)
	PortForwardServices(services)
}
