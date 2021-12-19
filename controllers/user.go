package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/nagokos/calorie_backend/models"
	"github.com/nagokos/calorie_backend/utils"
)

type JSONError struct {
	Error string `json:"error"`
	Field string `json:"field"`
	Code  int    `json:"code"`
}

func ApiError(w http.ResponseWriter, errMessage, field string, code int) {
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Field: field, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

func UserCreate(w http.ResponseWriter, req *http.Request) {
	utils.SetCors(w)

	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("\x1b[31m%s\x1b[0m\n", "Method Not Allowed")
		return
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}

	user := models.User{}

	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Println(err)
		return
	}

	if user.Name == "" {
		ApiError(w, "name is empty", "name", http.StatusUnprocessableEntity)
	}

	if user.Email == "" {
		ApiError(w, "email is empty", "email", http.StatusUnprocessableEntity)
	}

	if user.Password == "" {
		ApiError(w, "password is empty", "password", http.StatusUnprocessableEntity)
	}

	hashed := user.GenerateHash()
	err = user.Create(hashed)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				ApiError(w, "dluplicate email", "email", http.StatusUnprocessableEntity)
			}
		}
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
