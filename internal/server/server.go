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
		println("Get all works")

	case http.MethodPost:
		println("Post order")
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
func HandlerTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	switch r.Method {
	case http.MethodGet:
		task := database.GetTask(id)
		taskJson, err := json.Marshal(task)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(taskJson)
		println("Get works by task id")

	case http.MethodPost:
		println("Task post works by task id")
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
		println("Task delete by id")
		database.DeleteTask(id)

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
		/*orders,*/ tasks := database.GetOrder(id)
		/*ordersJson, err := json.Marshal(orders)
		if err != nil {
			log.Fatal(err)
		}*/
		tasksJson, err := json.Marshal(tasks)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		//w.Write(ordersJson)
		w.Write(tasksJson)
		println("Get works by order id")

	case http.MethodPut:

	case http.MethodPatch:

	case http.MethodDelete:
		println("Order delete by id")
		database.DeleteOrder(id)

	default:
		println("default")
	}

}

func Server() {
	fmt.Println("Server start...")

	router := mux.NewRouter()

	router.HandleFunc("/", HandlerAll)
	router.HandleFunc("/orders/{id:[0-9]+}", HandlerOrder)
	router.HandleFunc("/tasks/{id:[0-9]+}", HandlerTask)
	http.Handle("/", router)

	http.ListenAndServe(":8081", nil)
}
