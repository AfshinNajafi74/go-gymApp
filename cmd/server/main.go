package main

import (
	"log"
	"net/http"

	"github.com/AfshinNajafi74/go-gymApp/internal/config"
	"github.com/AfshinNajafi74/go-gymApp/pkg/database"
	"github.com/gorilla/mux"
)

func main() {

	// Load config
	cfg := config.Load()

	log.Println(cfg.DBUrl)

	// Connect database
	db := database.NewPostgres(cfg.DBUrl)
	_ = db

	// Router
	r := mux.NewRouter()

	r.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("gym app is running"))
	})

	log.Println("Server running on : " + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))

}
