package main

import (
	"encoding/json"
	"fmt"
	"golang-crud-sql/context"
	"golang-crud-sql/entity"
	"golang-crud-sql/handler"
	"golang-crud-sql/repository"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

func main() {

	//Retrieve data
	res, err := http.Get("https://random-data-api.com/api/users/random_user?size=10")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var posts []entity.UserRandom
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &posts)

	for p := range posts {
		fmt.Printf("%+v\n", posts[p])
	}

	// API
	db := context.Connect()
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	orderRepo := repository.NewOrderRepo(db)

	handler.UserRepo = userRepo
	userService := handler.NewUserHandler()

	handler.OrderRepo = orderRepo
	orderService := handler.NewOrderHandler()

	r := mux.NewRouter()
	r.HandleFunc("/users", userService.UserHandler)
	r.HandleFunc("/users/{id}", userService.UserHandler)

	r.HandleFunc("/orders", orderService.OrderHandler)
	r.HandleFunc("/orders/{id}", orderService.OrderHandler)

	fmt.Println("Now listening on port" + PORT)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
