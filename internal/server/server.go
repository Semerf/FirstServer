package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-redis/redis/v8"
	"github.com/semerf/FirstServer/internal/database"
	calcpb "github.com/semerf/FirstServer/proto"
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

	case http.MethodPut:

	case http.MethodPatch:

	case http.MethodDelete:
		println("Order delete by id")
		database.DeleteOrder(id)

	default:
		println("default")
	}

}

func HandlerResult(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	switch r.Method {
	case http.MethodGet:
		ctx := context.Background()
		var resJson []byte
		var res int32
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		val, err := rdb.Get(ctx, strconv.Itoa(id)).Result()

		if err == redis.Nil {
			tasks := database.GetOrder(id)
			conn, err := grpc.Dial("localhost:9997", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			var req calcpb.Request
			reqpt := make([](calcpb.Request), 0, len(tasks))
			reqspt := make([](*calcpb.Request), 0, len(tasks))

			for _, v := range tasks {
				req.TaskId = int32(v.Task_id)
				req.TaskName = v.Task_name
				req.Duration = int32(v.Duration)
				req.Resource = int32(v.Resource)
				req.PrevWork = v.Prev_work
				req.OrderId = int32(v.Order_id)
				reqpt = append(reqpt, req)
			}
			for i := range reqpt {
				reqspt = append(reqspt, &reqpt[i])
			}
			fmt.Println(tasks)
			fmt.Println(reqpt)
			client := calcpb.NewCalculatorClient(conn)
			var reqQ calcpb.RequestQuery
			reqQ.Requests = reqspt
			fmt.Println(reqQ)
			res = callCalc(client, &reqQ)

			resJson, err = json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			//
			err = rdb.Set(ctx, strconv.Itoa(id), res, time.Second*15).Err()
			if err != nil {
				panic(err)
			}

		} else if err != nil {
			panic(err)
		} else {
			vali, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			resJson, err = json.Marshal(vali)
			if err != nil {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		//w.Write(ordersJson)
		w.Write(resJson)

	case http.MethodPut:

	case http.MethodPatch:

	case http.MethodDelete:

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
	router.HandleFunc("/result/{id:[0-9]+}", HandlerResult)
	http.Handle("/", router)

	http.ListenAndServe(":8081", nil)
}
