package handler

import (
	"net/http"
	"storageproject/internal/domain"
	"storageproject/internal/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	service product.Service
}

func NewProduct(service product.Service) *Product {
	return &Product{
		service: service,
	}
}

func (m *Product) GetByName() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		product, err := m.service.GetByName(ctx, name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

func (m *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := m.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, products)
	}
}

func (m *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = m.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{"delete": id})
	}
}

func (m *Product) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var product domain.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		newId, errSave := m.service.Save(ctx, product)
		if errSave != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errSave.Error()})
			return
		}

		product, errGetId := m.service.GetById(ctx, newId)
		if errGetId != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errSave.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"product added ": product.Name})
	}
}

func (m *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{"error": "invalid ID"})
			return
		}
		var product domain.Product
		err = ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = m.service.Update(ctx, product, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"product updated": product.Name})
	}
}
