package main

import (
	"go-users/models"
	"go-users/routes"
)

func main() {
	models.Setup()
	routes.Setup()
}
