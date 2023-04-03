package main

import (
	"decomposition/app"
	"decomposition/repository"
)

func main() {
	repository := repository.NewMemStorage()
	app := app.New(repository)
	app.Run()
}
