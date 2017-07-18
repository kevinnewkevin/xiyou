package main

import (
	"fmt"
	"logic/application"
)

func main() {
	app := application.NewApp()
	app.Run()
	fmt.Println("Stop")
}
