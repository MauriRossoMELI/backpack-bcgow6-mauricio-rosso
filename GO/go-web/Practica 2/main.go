package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userGroup := router.Group("users")
	userGroup.POST("", CreateEntity)
	router.Run()
}

var usersStore []User //Para guardar las request que me van generando
var myAppToken string = "12345"

type User struct {
	Id         int       `json:"id"`
	Name       string    `json:"name" binding:"required"`
	Surname    string    `json:"surname" binding:"required"`
	Email      string    `json:"email" binding:"required"`
	Age        int       `json:"age" binding:"required"`
	Height     int       `json:"height" binding:"required"`
	IsActive   bool      `json:"isactive" binding:"required"`
	CreateDate time.Time `json:"createdate" binding:"required"`
}

func CreateEntity(ctx *gin.Context) {

	token := ctx.GetHeader("token")
	if token != myAppToken {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No tiene permisos para realizar la petición solicitada."})
		return
	}

	var userRequest User

	if err := ctx.ShouldBindJSON(&userRequest); err != nil { //GET ALL THE STRUCT USER BY BODY
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userRequest.Name == "" || userRequest.Surname == "" || userRequest.Email == "" || userRequest.Age == 0 || userRequest.Height == 0 || userRequest.CreateDate.IsZero() {
		if userRequest.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Field name is required!"})
		} else {
			if userRequest.Surname == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Field surname is required!"})
			} else {
				if userRequest.Email == "" {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": "Field email is required!"})
				} else {
					if userRequest.Age == 0 {
						ctx.JSON(http.StatusBadRequest, gin.H{"error": "Field age is required!"})
					} else {
						if userRequest.Height == 0 {
							ctx.JSON(http.StatusBadRequest, gin.H{"error": "Field 'height' is required!"})
						} else {
							if userRequest.CreateDate.IsZero() {
								ctx.JSON(http.StatusBadRequest, gin.H{"error": "Field 'createDate' is required!"})
							}
						}
					}
				}
			}
		}
		return
	}

	newId := GenerateId(&usersStore)

	newUser := User{
		Id:         newId,
		Name:       userRequest.Name,
		Surname:    userRequest.Surname,
		Email:      userRequest.Email,
		Age:        userRequest.Age,
		Height:     userRequest.Height,
		IsActive:   userRequest.IsActive,
		CreateDate: userRequest.CreateDate,
	}

	usersStore = append(usersStore, newUser)

	ctx.JSON(http.StatusOK, usersStore)
}

func GenerateId(usersTransactions *[]User) int {
	return len(*usersTransactions) + 1
}
