package main

import (
	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/cmd/server/handler"
	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/internal/users"
	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)

	service := users.NewService(repo)

	p := handler.NewUser(service)
	r := gin.Default()

	pr := r.Group("/users")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.UpdateSurnameAge())
	r.Run()
}
