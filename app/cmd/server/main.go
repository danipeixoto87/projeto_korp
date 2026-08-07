package main

import (
	"log"
	"net/http"

	"http-server-projeto-korp/internal/metrics"
	"http-server-projeto-korp/internal/routes"
)

func main() {

	metrics.Register()

	router := routes.NewRouter()

	log.Println("HTTP Server Projeto Korp iniciado na porta :8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}
}
