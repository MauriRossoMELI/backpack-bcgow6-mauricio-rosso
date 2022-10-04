package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crea un router con gin
	router := gin.Default()

	//EXERCISE 1.1
	// Captura la solicitud GET “/hello-world”
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Mauricio",
		})
	})

	//EXERCISE 1.2
	router.GET("/users", GetAll)

	//EXERCISE 2.1
	router.GET("/FilterUsers", GetUsers)

	//EXERCISE 2.2
	router.GET("/GetUserById", GetUserById)
	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}

// Handler filter id
func GetUserById(ctx *gin.Context) {
	jsonData, err := os.ReadFile("./users.json")
	if err != nil {
		panic("Error!")
	}
	users := &[]Users{}
	errJson := json.Unmarshal(jsonData, &users)
	if errJson != nil {
		panic(errJson.Error())
	}

	id := ctx.Request.URL.Query().Get("id") //GET THE USER ID BY URL
	idParsed, errParse := strconv.Atoi(id)

	if errParse != nil {
		panic("Error parsing ID requested!")
	}

	userResponse := Users{}

	for _, usr := range *users {
		if idParsed == usr.Id {
			userResponse.Id = usr.Id
			userResponse.Name = usr.Name
			userResponse.Surname = usr.Surname
			userResponse.Age = usr.Age
			userResponse.Email = usr.Email
			userResponse.Height = usr.Height
			userResponse.IsActive = usr.IsActive
			userResponse.CreateDate = usr.CreateDate
		}
	}

	if userResponse.Id == 0 {
		ctx.JSON(http.StatusNotFound, "No se encontró el usuario.")
	} else {
		ctx.JSON(http.StatusOK, userResponse)
	}
}

// Handler all filters
func GetUsers(ctx *gin.Context) {
	jsonData, err := os.ReadFile("./users.json")
	if err != nil {
		panic("Error!")
	}
	users := &[]Users{}
	errJson := json.Unmarshal(jsonData, &users)
	if errJson != nil {
		panic(errJson.Error())
	}

	var userRequest Users
	var usersResponse []Users

	if err := ctx.ShouldBindJSON(&userRequest); err != nil { //GET ALL THE STRUCT USER BY BODY
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, usr := range *users {
		if usr.Name == userRequest.Name || usr.Surname == userRequest.Surname || usr.Email == userRequest.Email || usr.Age == userRequest.Age || usr.Height == userRequest.Height || usr.IsActive == userRequest.IsActive {
			usersResponse = append(usersResponse, usr)
		}
	}
	ctx.JSON(http.StatusOK, usersResponse)
}

// Handler GetAll
func GetAll(ctx *gin.Context) {
	jsonData, err := os.ReadFile("./users.json")
	if err != nil {
		panic("Error!")
	}
	users := &[]Users{}
	errJson := json.Unmarshal(jsonData, &users)
	if errJson != nil {
		panic(errJson.Error())
	}
	ctx.JSON(http.StatusOK, users)
}

type Users struct {
	Id         int       `form:"id" json:"id"`
	Name       string    `form:"name" json:"name"`
	Surname    string    `form:"surname" json:"surname"`
	Email      string    `form:"email" json:"email"`
	Age        int       `form:"age" json:"age"`
	Height     int       `form:"height" json:"height"`
	IsActive   bool      `form:"isactive" json:"isactive"`
	CreateDate time.Time `form:"createdate" json:"createdate"`
}
