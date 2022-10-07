package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Database() {
	println("It's work!")

	//connStr := "user=posgres password=872000 dbname=Test sslmode=disable"
	connStr := "postgres://postgres:1234@localhost:5432/postgres"
	//eitherDB, err := pgx.Connect(context.Background(), connStr)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		panic(err)
	}
	result, err := db.Exec("SELECT * FROM testtable")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
