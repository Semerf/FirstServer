package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"time"
)

type Task struct {
	Task_name  string `json:"task_name"`
	Duration   byte   `json:"duration"`
	Resource   byte   `json:"resource"`
	Prev_tasks []Task `json:"previous_tasks"`
}

type Order struct {
	Order_Name string `json:"order_name"`
	Start_Date string `json:"start_date"`
	Tasks      []Task `json:"tasks"`
}

func main() {
	orders := []Order{
		{"Стройка", "18-09-2022", []Task{
			{"Задача 1", 2, 2, nil},
			{"Задача 2", 3, 1, nil},
		}},
	}
	fmt.Println(orders)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		switch r.Method {
		case http.MethodGet:
			productsJson, _ := json.Marshal(orders)
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(productsJson)
			break
		case http.MethodPost:
			decoder := json.NewDecoder(r.Body)
			task := Task{}
			err := decoder.Decode(&task)
			if err != nil {
				fmt.Print("Ошибка")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			break
		case http.MethodPut:
			break
		case http.MethodPatch:
			break
		case http.MethodDelete:
			break
		default:
			break
		}

	})

	http.ListenAndServe(":8081", nil)
}
