package main

import (
	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/cmd/server/handler"
	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)

	p := handler.NewUser(service)

	r := gin.Default()

	pr := r.Group("/users")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	r.Run()
}
