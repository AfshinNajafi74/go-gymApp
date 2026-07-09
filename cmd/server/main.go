package main

import (
	"log"
	"net/http"

	"github.com/AfshinNajafi74/go-gymApp/internal/config"
	"github.com/AfshinNajafi74/go-gymApp/internal/domain/profile"
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
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	// Load config
	cfg := config.Load()

	log.Println(cfg.DBUrl)

	// Connect database
	db := database.NewPostgres(cfg.DBUrl)
	err := db.AutoMigrate(
		&user.User{},
		&profile.AthleteProfile{},
		&profile.CoachProfile{},
	)

	if err != nil {
		log.Fatal(err)
	}

	// User
	userRepo := postgres.NewUserRepository(db)
	userService := user.NewService(userRepo)
	userHandler := userHttp.NewUserHandler(userService, cfg.JWTSecret)

	// Profile
	profileRepo := postgres.NewProfileRepository(db)
	profileService := profile.NewService(profileRepo)
	profileHandler := userHttp.NewProfileHandler(profileService)

	// Router
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// User routes
	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.Handle("/profile",
		userHttp.JWTMiddleware(cfg.JWTSecret)(http.HandlerFunc(userHandler.Profile)),
	).Methods("GET")

	jwt := userHttp.JWTMiddleware(cfg.JWTSecret)

	// Profile routes
	r.Handle("/profile/athlete",
		jwt(http.HandlerFunc(profileHandler.CreateAthleteProfile)),
	).Methods("POST")
	r.Handle("/profile/athlete",
		jwt(http.HandlerFunc(profileHandler.GetAthleteProfile)),
	).Methods("GET")
	r.Handle("/profile/athlete",
		jwt(http.HandlerFunc(profileHandler.UpdateAthleteProfile)),
	).Methods("PUT")

	r.Handle("/profile/coach",
		jwt(http.HandlerFunc(profileHandler.CreateCoachProfile)),
	).Methods("POST")
	r.Handle("/profile/coach",
		jwt(http.HandlerFunc(profileHandler.GetCoachProfile)),
	).Methods("GET")
	r.Handle("/profile/coach",
		jwt(http.HandlerFunc(profileHandler.UpdateCoachProfile)),
	).Methods("PUT")

	log.Println("Server running on : " + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))

}
