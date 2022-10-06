package main

import (
	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/cmd/server/handler"
	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/internal/users"
	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/pkg/store"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
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
