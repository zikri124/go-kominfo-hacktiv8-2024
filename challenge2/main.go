package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User
var lastUserIdCreated = 0

func main() {
	NetHttp()
}

func NetHttp() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// get all users
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users)
			return
		}

		// create user
		if r.Method == http.MethodPost {
			user := User{}
			// only bind username and email
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			lastUserIdCreated++
			user.ID = uint(lastUserIdCreated)
			users = append(users, user)

			w.WriteHeader(http.StatusAccepted)
			return
		}
	})

	// {id} => path variable
	http.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			pathValue, _ := strconv.Atoi(r.PathValue("id"))
			for _, user := range users {
				if user.ID == uint(pathValue) {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(user)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		}

		// mini quiz
		// buatlah method
		// PUT /users/:id untuk edit user by id
		// Delete /users/:id untuk delete user by id

	})

	// :8080 PORT
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Listen to port 3000")
	}
}
