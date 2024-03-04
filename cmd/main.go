package main

import (
	"github.com/gorilla/mux"
	"sldt-test-task/internal/server"
	"sldt-test-task/internal/service"
)

func main() {
	router := mux.NewRouter()

	verifySvc := service.NewVerficationService()
	srv := server.NewServer(router, verifySvc)

	srv.RunServer()
}
