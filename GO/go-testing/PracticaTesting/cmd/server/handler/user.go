package handler

import (
	"os"
	"strconv"

	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/internal/users"
	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/pkg/store/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	Age          int    `json:"age"`
	Height       int    `json:"height"`
	IsActive     bool   `json:"isactive"`
	CreationDate string `json:"creationdate"`
}

type User struct {
	service users.Service
}

func NewUser(p users.Service) *User {
	return &User{
		service: p,
	}
}

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
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

// StoreUsers godoc
// @Summary Store users
// @Tags Users
// @Description store users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param user body request true "User to store"
// @Success 200 {object} web.Response
// @Router /users [post]
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

// UpdateUsers godoc
// @Summary Update users
// @Tags Users
// @Description update users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id url int true "User id"
// @Param user body request true "User to update"
// @Success 200 {object} web.Response
// @Router /users [put]
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
		if req.CreationDate == "" {
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

// UpdateSurnameAge godoc
// @Summary Update users surname and age
// @Tags Users
// @Description update users surname and age
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id url int true "User id"
// @Param user body request true "Surname and age to update"
// @Success 200 {object} web.Response
// @Router /users [patch]
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

// DeleteUsers godoc
// @Summary Delete users
// @Tags Users
// @Description delete users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id url int true "User id to delete"
// @Success 200 {object} web.Response
// @Router /users [delete]
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
