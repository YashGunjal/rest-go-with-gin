package main

import (
	"rest-go-with-gin/models"

	"rest-go-with-gin/initialisers"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDB()
}

func main() {

	initialisers.DB.AutoMigrate(&models.POST{})
}

// Migrate the schema
