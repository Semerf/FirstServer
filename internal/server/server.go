package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/semerf/FirstServer/internal/database"
)

func HandlerAll(w http.ResponseWriter, r *http.Request) {
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
		println("Get works")

	case http.MethodPost:
		println("Post works")
		decoder := json.NewDecoder(r.Body)
		order := database.Order{}
		err := decoder.Decode(&order)
		if err != nil {
			fmt.Print("Ошибка")
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(order)
		database.AddOrder(order)

	case http.MethodPut:

	case http.MethodPatch:

	case http.MethodDelete:

	default:
		println("default")
	}

}
func HandlerOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	switch r.Method {
	case http.MethodGet:
		orders, tasks := database.GetOrder(id)
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
		println("Get works by order id")

	case http.MethodPost:
		println("Task post works by order id")
		decoder := json.NewDecoder(r.Body)
		task := database.Task{}
		err := decoder.Decode(&task)
		if err != nil {
			fmt.Print("Ошибка")
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(task)
		database.AddTask(task, id)

	case http.MethodPut:

	case http.MethodPatch:

	case http.MethodDelete:

	default:
		println("default")
	}

}

func Server() {
	fmt.Println("It's work too...")

	router := mux.NewRouter()

	router.HandleFunc("/", HandlerAll)
	router.HandleFunc("/{id:[0-9]+}", HandlerOrder)
	http.Handle("/", router)

	http.ListenAndServe(":8081", nil)
}
