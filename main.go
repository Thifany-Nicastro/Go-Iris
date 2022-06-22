package main

import "go-iris/config"

func main() {
	app := config.NewApp()

	app.Listen("localhost:8080")
}
