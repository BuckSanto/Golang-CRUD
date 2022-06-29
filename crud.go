package main

import (
	"fmt"
	"golang-crud/handler"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

func main() {
	r := mux.NewRouter()

	handler := handler.NewUserHandler()

	r.HandleFunc("/users", handler.UserHandler)
	r.HandleFunc("/users/{id}", handler.UserHandler)

	fmt.Println("Now listening on port" + PORT)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
