package handler

import (
	"internal/users"
	"time"
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
