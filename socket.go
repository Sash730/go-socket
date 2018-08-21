package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Sash730/go-socket/controller"
)



func main() {


	c := controller.NewRecentlyController()

	r := mux.NewRouter()

	// HEALTH CHECK
	r.HandleFunc("/socket.io", c.ViewReport)

	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
