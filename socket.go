package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Sash730/go-socket/controller"
	"github.com/gorilla/handlers"
	"os"
)



func main() {


	c := controller.NewRecentlyController()

	r := mux.NewRouter()

	// HEALTH CHECK
	r.HandleFunc("/socket.io", c.ViewReport)

	// these two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	router := handlers.LoggingHandler(os.Stdout, r)
	router = handlers.CORS(allowedOrigins, allowedMethods)(router)

	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
