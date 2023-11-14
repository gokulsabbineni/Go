package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	api "myproject/api/library"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:Gokul@26@tcp(127.0.0.1:3306)/sys?parseTime=true"

	fmt.Println("My Connection", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("DB Drivers established")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection Established sucessfully")
	api.RegisterRoutes(db)
	fmt.Println("Sent the request to routers")

	log.Println("server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
