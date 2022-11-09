package routes

import (
	"database/sql"
	"storageproject/cmd/server/handler"
	"storageproject/internal/product"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)
	r.rg.GET("/products", handler.GetAll())
	r.rg.GET("/products/:name", handler.GetByName())
	r.rg.DELETE("/products/:id", handler.Delete())
	// r.rg.POST("/movies", handler.Create())
	// r.rg.PATCH("/movies/:id", handler.Update())
}
