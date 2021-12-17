package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nagokos/calorie_backend/config"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func setCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

func createUser(w http.ResponseWriter, req *http.Request) {
	setCors(w)

	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("\x1b[31m%s\x1b[0m\n", "Method Not Allowed")
		return
	}

	buffer, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	b := User{}

	if err := json.Unmarshal(buffer, &b); err != nil {
		fmt.Println("err")
	}

	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/api/v1/users", createUser)
	log.Print(http.ListenAndServe(":"+config.Config.Port, nil))
}
