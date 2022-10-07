package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Order struct {
	order_id   int
	order_name string
	start_date string
}
type Task struct {
	task_id   int
	task_name string
	duration  string
	resource  int
	prev_work []int
	order_id  int
}

func Database() {
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
		err := OrderRows.Scan(&ordr.order_id, &ordr.order_name, &ordr.start_date)
		if err != nil {
			log.Fatal(err)
		}
		ordrs = append(ordrs, ordr)
	}
	for _, ordr := range ordrs {
		fmt.Printf("%d, %s, %s \n", ordr.order_id, ordr.order_name, ordr.start_date)
	}
	fmt.Println("Записи из таблици tasks...")
	TaskRows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	tsks := make([]*Task, 0)
	for TaskRows.Next() {
		tsk := new(Task)
		err := TaskRows.Scan(&tsk.task_id, &tsk.task_name, &tsk.duration, &tsk.resource, &tsk.prev_work, &tsk.order_id)
		if err != nil {
			log.Fatal(err)
		}
		tsks = append(tsks, tsk)
	}
	for _, tsk := range tsks {
		fmt.Printf("%d, %s, %s, %d, ", tsk.task_id, tsk.task_name, tsk.duration, tsk.resource)
		fmt.Print(tsk.prev_work)
		fmt.Printf("%d \n", tsk.order_id)
	}

	//fmt.Println(ordrs)
}
