package main

import (
	"log"
	"net/http"
	"task-api/src/config"
	"task-api/src/infrastructure/db"
	"task-api/src/infrastructure/router"
)

func main() {
	conf := config.Get()

	sqlhandler := db.NewSqlhandler()
	defer sqlhandler.Close()

	r := router.Handler(sqlhandler)
	log.Fatal(http.ListenAndServe(":"+conf.Server.Port, r))
}
