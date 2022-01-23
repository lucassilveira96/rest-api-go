package main

import (
	"fmt"
	"rest-api-go/database"
	"rest-api-go/routes"
)

func main() {
	database.DBConnection()
	fmt.Println("Iniciando o servidor rest go")
	routes.HandleRequest()
}
