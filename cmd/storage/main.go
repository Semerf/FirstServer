package main

import (
	"fmt"

	"github.com/semerf/FirstServer/internal/database"
	"github.com/semerf/FirstServer/internal/server"
)

func main() {
	fmt.Println("It's work!")
	database.DatabaseShow()

	fmt.Println(database.GetDatabase())
	server.Server()

}
