package main

import "github.com/christiangda/mango/api"

func main() {
	var name = "MyAPI"
	var version = "1.0.0"
	var port = "8080"
	var environment = "development"

	myapi := api.NewAPI(name, version, port, environment)

	myapi.Run()
}
