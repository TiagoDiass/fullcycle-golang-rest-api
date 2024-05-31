package main

import (
	"net/http"

	"github.com/TiagoDiass/fullcycle-golang-rest-api/configs"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/infra/database"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProductDB(db)
	userDB := database.NewUserDB(db)
	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB, cfg.TokenAuth, cfg.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.ListProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/{id}", productHandler.GetProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/session", userHandler.CreateSession)

	http.ListenAndServe(":8000", r)
}
