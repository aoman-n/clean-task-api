package main

import (
	"log"
	"net/http"
	"task-api/src/config"
	"task-api/src/infrastructure/db"
	"task-api/src/infrastructure/router"

	"github.com/go-playground/validator/v10"
)

func main() {
	conf := config.Get()

	sqlhandler := db.NewSqlhandler()
	defer sqlhandler.Close()
	validator := validator.New()

	r := router.Handler(sqlhandler, validator)
	log.Fatal(http.ListenAndServe(":"+conf.Server.Port, r))
}
