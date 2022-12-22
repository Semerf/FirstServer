package main

import (
	"fmt"

	"github.com/semerf/FirstServer/internal/database"
	"github.com/semerf/FirstServer/internal/server"
)

func main() {

	var choice int
	go server.Server()

	fmt.Println("Список orders и tasks")
	database.DatabaseShow()
	fmt.Println("Выберите необходимый order")
	fmt.Scan(&choice)

}
