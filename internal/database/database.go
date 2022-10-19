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
	Duration  int    `json:"duration"`
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
		fmt.Printf("%d, %s, %d, %d, %s, %d \n", tsk.Task_id, tsk.Task_name, tsk.Duration, tsk.Resource, tsk.Prev_work, tsk.Order_id)
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

func GetOrder(id int) /*[]Order,*/ []Task {
	connStr := "postgres://postgres:1234@localhost:5432/cslab"
	//eitherDB, err := pgx.Connect(context.Background(), connStr)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	/*OrderRows, err := db.Query("SELECT * FROM orders WHERE order_id = $1", id)
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
	}*/
	TaskRows, err := db.Query("SELECT * FROM tasks WHERE order_id = $1", id)
	if err != nil {
		log.Fatal(err)
		return nil
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
	return /*ordrs,*/ tsks
}

func GetTask(id int) Task {
	connStr := "postgres://postgres:1234@localhost:5432/cslab"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	TaskRows, err := db.Query("SELECT * FROM tasks WHERE task_id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	tsk := *new(Task)
	for TaskRows.Next() {
		err := TaskRows.Scan(&tsk.Task_id, &tsk.Task_name, &tsk.Duration, &tsk.Resource, &tsk.Prev_work, &tsk.Order_id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return tsk
}

func AddOrder(order Order) {
	connStr := "postgres://postgres:1234@localhost:5432/cslab"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	sqlStatement := "INSERT INTO orders (order_name, start_date) VALUES ($1, $2)"

	_, err = db.Exec(sqlStatement, order.Order_name, order.Start_date)
	if err != nil {
		log.Fatal(err)
	}
}

func AddTask(task Task, id int) {
	connStr := "postgres://postgres:1234@localhost:5432/cslab"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	sqlStatement := "INSERT INTO tasks (task_name, duration, resource, prev_tasks, order_id) VALUES ($1, $2, $3, $4, $5)"

	_, err = db.Exec(sqlStatement, task.Task_name, task.Duration, task.Resource, task.Prev_work, id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTask(id int) {
	connStr := "postgres://postgres:1234@localhost:5432/cslab"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	sqlStatement := "DELETE FROM tasks WHERE task_id = $1"

	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteOrder(id int) {
	connStr := "postgres://postgres:1234@localhost:5432/cslab"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	sqlStatementTask := "DELETE FROM tasks WHERE order_id = $1"

	_, err = db.Exec(sqlStatementTask, id)
	if err != nil {
		log.Fatal(err)
	}
	sqlStatementOrder := "DELETE FROM orders WHERE order_id = $1"

	_, err = db.Exec(sqlStatementOrder, id)
	if err != nil {
		log.Fatal(err)
	}
}
