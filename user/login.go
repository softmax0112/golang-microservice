// Package Classification for User API
//
//Documentation for Login API
//
// Shemes: http
// BasePath: /login
// Version 1.0.0
// Host: localhost

// Consumes:
// - applcation/json

// Produces:
// - application/json
// swagger :meta

package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// swagger:route /login
// Show details of Student
// responses:
//   200:

type user struct {
	Username string `schema:"username"`
	Password string `schema:"password"`
	Marks    []int  `json:"marks"`
}

//Login Route
func (u DBConfig) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := u.DB

	// Realtime checking that DB connection is alive or not
	err := db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"errorCode": err.Error()})
		return
	}
	errorMap := make(map[string]string)
	var user user
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &user)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	username := strings.TrimSpace(user.Username)
	password := strings.TrimSpace(user.Password)
	if len(username) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		errorMap["errorCode"] = "username must not be empty"
		json.NewEncoder(w).Encode(errorMap)
		return
	}
	if len(password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		errorMap["errorCode"] = "password must not be empty"
		json.NewEncoder(w).Encode(errorMap)
		return
	}
	var marks *[]byte
	err = db.QueryRow("SELECT marks FROM users WHERE username=? AND password=?", username, password).Scan(&marks)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMap["errorCode"] = "UserNotFound"
		errorMap["description"] = err.Error()
		json.NewEncoder(w).Encode(errorMap)
		return
	}
	err = json.Unmarshal(*marks, &user.Marks)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMap["errorCode"] = "Error in Marks Unmarshalling"
		errorMap["description"] = err.Error()
		json.NewEncoder(w).Encode(errorMap)
		return
	}
	response := make(map[string][]int)
	response["marks"] = user.Marks
	json.NewEncoder(w).Encode(response)

}
