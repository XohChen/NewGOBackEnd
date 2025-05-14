package main

import (
	"log"

	"github.com/XohChen/NewGOBackEnd/cmd/api"
	"github.com/XohChen/NewGOBackEnd/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "asd",
		Addr:                 "127.0.0.1:3306",
		DBName:               "test_go",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	server := api.NewAPIServer(":8081", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
