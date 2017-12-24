package main

import "github.com/christiangda/mango/api"

func main() {
	var name = "MyAPI"
	var version = "1.0.0"
	var port = "8080"

	myapi, _ := api.NewAPI(name, version, port)
	myapi.Run()
}
