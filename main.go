package main

import (
	"bms/config"
	"bms/controller"
)

func main() {
	config.EstablishDatabaseConnection("root:admin@tcp(127.0.0.1:3306)/bms?charset=utf8mb4&parseTime=True&loc=Local")
	server := controller.NewEchoServer()
	defer server.Close()
	err := server.Start("3306")
	if err != nil {
		panic(err)
	}
}

func init() {
	/*	err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		url := .Database.URL
		port := config.Server.Port
		fmt.Println("DB URL:", url)
		fmt.Println("Server Port:", port)*/
}
