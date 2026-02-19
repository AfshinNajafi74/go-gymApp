package main

import (
	"log"
	"net/http"

	"github.com/AfshinNajafi74/go-gymApp/internal/config"
	"github.com/AfshinNajafi74/go-gymApp/internal/domain/user"
	userHttp "github.com/AfshinNajafi74/go-gymApp/internal/handler/http"
	"github.com/AfshinNajafi74/go-gymApp/internal/repository/postgres"
	"github.com/AfshinNajafi74/go-gymApp/pkg/database"
	"github.com/gorilla/mux"

	_ "github.com/AfshinNajafi74/go-gymApp/docs"
	"github.com/swaggo/http-swagger"
)

// @title Gym App API
// @version 1.0
// @description Backend API for Gym App
// @host localhost:8080
// @BasePath /
func main() {

	// Load config
	cfg := config.Load()

	log.Println(cfg.DBUrl)

	// Connect database
	db := database.NewPostgres(cfg.DBUrl)
	_ = db

	userRepo := postgres.NewUserRepository(db)

	userService := user.NewService(userRepo)
	userHandler := userHttp.NewUserHandler(userService)

	// Router
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	log.Println("Server running on : " + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))

}
