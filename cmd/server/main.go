package main

import (
	"net/http"

	"github.com/TiagoDiass/fullcycle-golang-rest-api/configs"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/infra/database"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.ListProducts)

	r.Post("/users", userHandler.CreateUser)
	r.Post("/session", userHandler.CreateSession)

	http.ListenAndServe(":8000", r)
}
