package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
POST /v1/user/details to create a user
GET /v1/user/details to retrieve user information (use mock data)
*/
var users = make([]User, 0)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

func getUsersDataHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(&users)
	if err != nil {
		log.Fatalln("There was some error ecoding user to writer")
	}
}

func createUserHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Fatalln("There was some error decoding the request body to struct user")
	}

	users = append(users, user)

	if err := json.NewEncoder(writer).Encode(&user); err != nil {
		log.Fatalln("There was some error ecoding user to writer")
	}

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/user/details", getUsersDataHandler).Methods("GET")
	r.HandleFunc("/v1/user/details", createUserHandler).Methods("POST")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home Page")
	})
	fmt.Println("Server started on Port 8080")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalln("Server is unable to start")
	}

}
