package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"test/config"
	"test/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/pdrum/swagger-automation/docs"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(" Error Occured \n", err.Error())
	}

	db, err := sql.Open(config.DBType, config.DBUserName+":"+config.DBPassword+"@tcp("+config.DBHost+":"+config.DBPort+")/"+config.DBName)
	defer db.Close()
	if err != nil {
		log.Fatal("Problem in Database connection, " + err.Error())
	}
	userHandler := user.DBConfig{DB: db}

	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/user/login", userHandler.Login).Methods("GET")
	router.HandleFunc("/user/register", userHandler.Register).Methods("GET")
	fmt.Println("Server Started at " + config.ApplicationPort)
	log.Fatal(http.ListenAndServe(":"+config.ApplicationPort, router))
}
