package main

import (
	// controller "ConnectApp/src/controllers"
	"ConnectApp/src/controllers"
	"ConnectApp/src/middlewares"
	routes "ConnectApp/src/routes"
	db "ConnectApp/src/sql"

	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const port = 8080

func main() {
	err := db.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")
	log.Println("listening on port", port)
	go controllers.ListenToWsChannel()
	// Consume and print values from the channel
	errors := http.ListenAndServe(fmt.Sprintf(":%d", port), middlewares.EnableCORS(routes.Routes()))
	// errors := http.ListenAndServe(fmt.Sprintf(":%d", port), routes.Routes())
	if errors != nil {
		log.Fatal(errors)
	}

}
