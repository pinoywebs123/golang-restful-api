package main

import (
	"example.com/crud/initializers"
	"example.com/crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
