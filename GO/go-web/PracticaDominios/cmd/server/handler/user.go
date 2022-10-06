package handler

import (
	"os"
	"strconv"
	"time"

	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/internal/users"
	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/pkg/store/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	Height       int       `json:"height"`
	IsActive     bool      `json:"isactive"`
	CreationDate time.Time `json:"creationdate"`
}

type User struct {
	service users.Service
}

func NewUser(p users.Service) *User {
	return &User{
		service: p,
	}
}

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		users, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		if len(users) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "No users found."))
			return
		}

		ctx.JSON(200, web.NewResponse(200, users, ""))
	}
}

func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		user, err := c.service.Store(req.Name, req.Surname, req.Email, req.Age, req.Height, req.IsActive, req.CreationDate)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, user, "User stored successfully."))
	}
}

func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID!"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "The name is required!"))
			return
		}
		if req.Surname == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "The surname is required!"))
			return
		}
		if req.Email == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "The email is required!"))
			return
		}
		if req.Age == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "The age is required!"))
			return
		}
		if req.Height == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "The height is required!"))
			return
		}
		if req.CreationDate.IsZero() {
			ctx.JSON(400, web.NewResponse(400, nil, "The creation date is required!"))
			return
		}
		user, err := c.service.Update(int(id), req.Name, req.Surname, req.Email, req.Age, req.Height, req.IsActive, req.CreationDate)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, user, "User updated successfully!"))
	}
}

func (c *User) UpdateSurnameAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID!"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Surname == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "The surname is required!"))
			return
		}
		if req.Age == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "The age is required!"))
			return
		}
		user, err := c.service.UpdateSurnameAge(int(id), req.Surname, req.Age)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, user, "User updated (surname & age) successfully!"))
	}
}

func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID!"))
			return
		}
		errId := c.service.Delete(id)
		if errId != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, nil, "User deleted successfully!"))
	}
}
