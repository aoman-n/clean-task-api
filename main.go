package main

import (
	"log"
	"net/http"
	"task-api/src/config"
	"task-api/src/infrastructure"
)

func main() {
	conf := config.Get()

	sqlhandler := infrastructure.NewSqlhandler()
	defer sqlhandler.Close()

	router := infrastructure.Handler()
	log.Fatal(http.ListenAndServe(":"+conf.Server.Port, router))
}
