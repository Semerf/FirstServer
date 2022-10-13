package main

import (
	"fmt"

	"github.com/semerf/FirstServer/internal/database"
	"github.com/semerf/FirstServer/internal/server"
)

func main() {

	orders, tasks := database.GetDatabase()
	fmt.Println(orders)
	fmt.Println(tasks)

	server.Server()
}
