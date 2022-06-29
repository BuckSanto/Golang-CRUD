package handler

import (
	"encoding/json"
	"fmt"
	"golang-crud/entity"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var users = map[int]entity.User{
	1: {
		Id:        1,
		Username:  "budi123",
		Email:     "budi123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()},
	2: {
		Id:        2,
		Username:  "cantya123",
		Email:     "cantya123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()},
	3: {
		Id:        3,
		Username:  "cantya123",
		Email:     "dantya123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()},
	4: {
		Id:        4,
		Username:  "cantya123",
		Email:     "dantya123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()},
	5: {
		Id:        5,
		Username:  "andi123",
		Email:     "andi123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()},
}

type UserHandlerIface interface {
	UserHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
}

func NewUserHandler() UserHandlerIface {
	return &UserHandler{}
}

func (u *UserHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodGet:
		if id != "" {
			getUserById(w, r, id)
		} else {
			getUsers(w, r)
		}
	case http.MethodPost:
		registerUser(w, r)
	case http.MethodPut:
		updateUser(w, r, id)
	case http.MethodDelete:
		deleteUser(w, r, id)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var usersSlice []entity.User
	for _, v := range users {
		usersSlice = append(usersSlice, v)
	}
	json, _ := json.Marshal(usersSlice)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func getUserById(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			if user, ok := users[idInt]; ok {
				jsonData, _ := json.Marshal(user)
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
				return
			} else {
				w.Write([]byte("User not found"))
				return
			}
		}
	}
	w.Write([]byte("Invalid parameter"))
	return
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user entity.User
	w.Header().Add("Content-Type", "application/json")

	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("Error decoding json body"))
		return
	}

	if _, found := users[user.Id]; found {
		w.Write([]byte("User already exist"))
		return
	}

	users[user.Id] = user
	var usersSlice []entity.User
	for _, v := range users {
		usersSlice = append(usersSlice, v)
	}
	jsonData, _ := json.Marshal(&usersSlice)
	w.Write((jsonData))
}

func updateUser(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			if _, found := users[idInt]; found {
				decoder := json.NewDecoder(r.Body)
				var userSlice entity.User
				if err := decoder.Decode(&userSlice); err != nil {
					w.Write([]byte("Error decoding json body"))
					return
				}
				users[idInt] = userSlice
				jsonData, _ := json.Marshal(&userSlice)
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
				return
			} else {
				w.Write([]byte("User not found"))
				return
			}
		}
	}
	w.Write([]byte("Invalid parameter"))
	return
}

func deleteUser(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			if _, found := users[idInt]; found {
				delete(users, idInt)
				w.Write([]byte(fmt.Sprintf("User %d deleted", idInt)))
				return
			} else {
				w.Write([]byte("User not found"))
				return
			}
		}
	}
	w.Write([]byte("Invalid parameter"))
	return
}
