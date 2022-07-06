package handler

import (
	"context"
	"encoding/json"
	"golang-crud-sql/entity"
	"golang-crud-sql/repository"
	"net/http"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

var UserRepo repository.UserRepoIface

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

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	users, err := UserRepo.GetUsers(ctx)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func getUserById(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" {
		if idInt, err := strconv.Atoi(id); err == nil {

			ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelfunc()
			user, err := UserRepo.GetUserById(ctx, idInt)

			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			if user.Id != 0 {
				jsonData, _ := json.Marshal(user)
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
				return
			}
			w.Write([]byte("User not found"))
			return
		}
	}
	w.Write([]byte("Invalid parameter"))
	return
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user entity.User

	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("Error decoding json body"))
		return
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	result, err := UserRepo.CreateUser(ctx, user)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(result))
}

func updateUser(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			decoder := json.NewDecoder(r.Body)
			var userSlice entity.User
			if err := decoder.Decode(&userSlice); err != nil {
				w.Write([]byte("Error decoding json body"))
				return
			}

			ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelfunc()
			result, err := UserRepo.UpdateUser(ctx, idInt, userSlice)

			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.Write([]byte(result))
			return
		}
	}
	w.Write([]byte("Invalid parameter"))
	return
}

func deleteUser(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" {
		if idInt, err := strconv.Atoi(id); err == nil {

			ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelfunc()
			result, err := UserRepo.DeleteUser(ctx, idInt)

			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.Write([]byte(result))
			return
		}
	}
	w.Write([]byte("Invalid parameter"))
	return
}
