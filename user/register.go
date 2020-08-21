package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/schema"
)

type regForm struct {
	Username string `schema:"username,required"`
	Password string `schema:"password,required"`
	Marks    []int  `schema:"marks"`
}

//Register function is for regitration route
func (d DBConfig) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := d.DB
	err := db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"errorCode": "Problem while checking Connection/Ping", "description": err.Error()})
		return
	}
	var formData regForm
	// Decoding into Struct from Query Parameters by Gorilla's Schema package
	err = schema.NewDecoder().Decode(&formData, r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"errorCode": err.Error()})
		return
	}

	stmt, err := db.Prepare("INSERT INTO users (username,password,marks) VALUES (?,?,?)")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"description": err.Error(), "errorCode": "Unable to get a connection"})
		return
	}

	result, err := stmt.Exec(formData.Username, formData.Password, strings.Replace(fmt.Sprintf("%v", formData.Marks), " ", ",", -1))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"description": err.Error(), "errorCode": "User Exists"})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"errorCode": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]int64{"Rows Affected": rowsAffected})

}
