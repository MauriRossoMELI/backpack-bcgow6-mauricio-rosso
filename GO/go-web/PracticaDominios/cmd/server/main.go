package main

import (
	"os"

	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/cmd/server/handler"
	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/internal/users"
	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/pkg/store"

	"github.com/joho/godotenv"

	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)

	service := users.NewService(repo)

	p := handler.NewUser(service)
	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/users")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.UpdateSurnameAge())
	r.Run()
}
