package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/larissavarjao/sintomas-api/api/pacient"
	"github.com/larissavarjao/sintomas-api/api/symptom"
	"github.com/larissavarjao/sintomas-api/api/user"
	"github.com/larissavarjao/sintomas-api/handlers"
	"github.com/larissavarjao/sintomas-api/middlewares"
	_ "github.com/lib/pq"
)

const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "password"
const DbName = "sintomas-develop"

var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

func main() {
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	userService := user.NewService(db)
	pacientService := pacient.NewService(db)
	symptomService := symptom.NewService(db)

	r := mux.NewRouter()
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(middlewares.ValidateAuthentication())

	handlers.UsersHandlers(r, n, userService)
	handlers.PacientsHandlers(r, n, pacientService)
	handlers.SymptomsHandlers(r, n, symptomService)

	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr: ":8080",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}