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
		if r.Method == http.MethodPut {
			userTemp := User{}
			if err := json.NewDecoder(r.Body).Decode(&userTemp); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			pathValue, _ := strconv.Atoi(r.PathValue("id"))

			for i, user := range users {
				if user.ID == uint(pathValue) {
					user.Username = userTemp.Username
					user.Email = userTemp.Email

					users[i] = user

					w.WriteHeader(http.StatusAccepted)
					json.NewEncoder(w).Encode(users[i])
					return
				}
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(User{})
			return
		}

		// Delete /users/:id untuk delete user by id
		if r.Method == http.MethodDelete {
			pathValue, _ := strconv.Atoi(r.PathValue("id"))

			for i, user := range users {
				if user.ID == uint(pathValue) {
					users[i] = user

					users = append(users[0:i], users[i+1:]...)

					w.WriteHeader(http.StatusAccepted)
					json.NewEncoder(w).Encode(users)
					return
				}
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Success : false")
			return
		}
	})

	// :8080 PORT
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Listen to port 80")
	}
}
