package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qobilovvv/patterns/service_repo/handler"
	"github.com/qobilovvv/patterns/service_repo/models"
	"github.com/qobilovvv/patterns/service_repo/repository"
	"github.com/qobilovvv/patterns/service_repo/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	repo := repository.NewUserRepositoryDB(db)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	r.Get("/users", h.GetAllUsers)
	r.Post("/user/create", h.CreateUser)

	fmt.Println("Server started at :3000")
	http.ListenAndServe(":3000", r)
}
