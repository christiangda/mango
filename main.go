package mango

import "github.com/christiangda/mango/api"

func main() {
	var name = "MyAPI"
	var version = "1.0.0"
	var port = "8080"

	api := api.NewAPI(name, version, port)
	api.Run()
}
