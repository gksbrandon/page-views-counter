package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gksbrandon/page-views-counter/driver"

	"github.com/gksbrandon/page-views-counter/controllers"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

var db *sql.DB

// Initializes environment variables
func init() {
	gotenv.Load()
}

// Error checking
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Connect db instance
	db = driver.ConnectDB()

	// Sets controller methods
	controller := controllers.Controller{}

	// Initialize router, endpoints and handler functions
	router := mux.NewRouter()
	router.HandleFunc("/counter/v1/statistics/article_id/{article_id}", controller.GetView(db)).Methods("GET")
	router.HandleFunc("/counter/v1/statistics", controller.AddView(db)).Methods("POST")

	// Run server
	fmt.Println("Server is running at port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
