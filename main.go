package main

import (
	"decomposition/app"
	"decomposition/repository"
)

func main() {
	storage := repository.NewMemStorage()
	instanceApp := app.New(storage)
	instanceApp.Run()
}
