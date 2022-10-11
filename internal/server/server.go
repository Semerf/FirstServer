package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/semerf/FirstServer/internal/database"
)

func Server() {
	fmt.Println("It's work too...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		switch r.Method {
		case http.MethodGet:
			orders, tasks := database.GetDatabase()
			ordersJson, err := json.Marshal(orders)
			if err != nil {
				log.Fatal(err)
			}
			tasksJson, err := json.Marshal(tasks)
			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(ordersJson)
			w.Write(tasksJson)
			println("all works")

		case http.MethodPost:
			decoder := json.NewDecoder(r.Body)
			task := database.Task{}
			err := decoder.Decode(&task)
			if err != nil {
				fmt.Print("Ошибка")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		case http.MethodPut:

		case http.MethodPatch:

		case http.MethodDelete:

		default:
			println("default")
		}

	})

	http.ListenAndServe(":8081", nil)
}
