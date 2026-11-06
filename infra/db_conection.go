package infra

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getConnection() string {
	return "user=postgres host=localhost port=5432 password=2173 dbname=payreminder sslmode=disable"
}


func DBConntion() (*sqlx.DB, error) {

	dbSource := getConnection()
	dbCon, err := sqlx.Connect("postgres", dbSource)

	if err != nil{
		fmt.Println(err)
		return  nil, err
	}

	return  dbCon, nil
}