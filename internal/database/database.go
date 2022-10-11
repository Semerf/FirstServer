package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Order struct {
	Order_id   int    `json:"order_id"`
	Order_name string `json:"order_name"`
	Start_date string `json:"start_date"`
}
type Task struct {
	Task_id   int    `json:"task_id"`
	Task_name string `json:"task_name"`
	Duration  string `json:"duration"`
	Resource  int    `json:"resource"`
	Prev_work string `json:"previous_tasks"`
	Order_id  int    `json:"order_id"`
}

func DatabaseShow() {
	println("It's work!")

	//connStr := "user=posgres password=872000 dbname=Test sslmode=disable"
	connStr := "postgres://postgres:1234@localhost:5432/cslab"
	//eitherDB, err := pgx.Connect(context.Background(), connStr)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Записи из таблици orders...")
	OrderRows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Fatal(err)
	}
	ordrs := make([]*Order, 0)
	for OrderRows.Next() {
		ordr := new(Order)
		err := OrderRows.Scan(&ordr.Order_id, &ordr.Order_name, &ordr.Start_date)
		if err != nil {
			log.Fatal(err)
		}
		ordrs = append(ordrs, ordr)
	}
	for _, ordr := range ordrs {
		fmt.Printf("%d, %s, %s \n", ordr.Order_id, ordr.Order_name, ordr.Start_date)
	}
	fmt.Println("Записи из таблици tasks...")
	TaskRows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	tsks := make([]*Task, 0)
	for TaskRows.Next() {
		tsk := new(Task)
		err := TaskRows.Scan(&tsk.Task_id, &tsk.Task_name, &tsk.Duration, &tsk.Resource, &tsk.Prev_work, &tsk.Order_id)
		if err != nil {
			log.Fatal(err)
		}
		tsks = append(tsks, tsk)
	}
	for _, tsk := range tsks {
		fmt.Printf("%d, %s, %s, %d, %s, %d \n", tsk.Task_id, tsk.Task_name, tsk.Duration, tsk.Resource, tsk.Prev_work, tsk.Order_id)
	}
}

func GetDatabase() ([]Order, []Task) {
	connStr := "postgres://postgres:1234@localhost:5432/cslab"
	//eitherDB, err := pgx.Connect(context.Background(), connStr)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	OrderRows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Fatal(err)
	}
	ordrs := make([]Order, 0)
	for OrderRows.Next() {
		ordr := *new(Order)
		err := OrderRows.Scan(&ordr.Order_id, &ordr.Order_name, &ordr.Start_date)
		if err != nil {
			log.Fatal(err)
		}
		ordrs = append(ordrs, ordr)
	}
	TaskRows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	tsks := make([]Task, 0)
	for TaskRows.Next() {
		tsk := *new(Task)
		err := TaskRows.Scan(&tsk.Task_id, &tsk.Task_name, &tsk.Duration, &tsk.Resource, &tsk.Prev_work, &tsk.Order_id)
		if err != nil {
			log.Fatal(err)
		}
		tsks = append(tsks, tsk)
	}
	return ordrs, tsks
}
