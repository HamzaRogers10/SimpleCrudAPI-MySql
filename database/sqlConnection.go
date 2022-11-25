package database

import (
	"database/sql"
	"fmt"
	"log"
)

func Sqlclient() *sql.DB {
	fmt.Println("Welcome To SQL connection")
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Hamza@10"
	dbName := "lost_listing"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=True")
	if err != nil {
		log.Fatal("error connecting DB : ", err.Error())
	}
	fmt.Println("Connection Successfully")

	return db

}
