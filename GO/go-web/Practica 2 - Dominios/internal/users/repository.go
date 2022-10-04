package users

import (
	"fmt"
	"time"
)

type User struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	Height       int       `json:"height"`
	IsActive     bool      `json:"isactive"`
	CreationDate time.Time `json:"creationdate"`
}

var ps []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name string, surname string, email string, age int, height int, isActive bool, creationDate time.Time) (User, error)
	LastID() (int, error)
	Update(id int, name string, surname string, email string, age int, height int, isActive bool, creationDate time.Time) (User, error)
}

type repository struct{} //struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, name string, surname string, email string, age int, height int, isActive bool, creationDate time.Time) (User, error) {
	p := User{id, name, surname, email, age, height, isActive, creationDate}
	ps = append(ps, p)
	lastID = p.Id
	return p, nil
}

func (r *repository) Update(id int, name string, surname string, email string, age int, height int, isActive bool, creationDate time.Time) (User, error) {
	p := User{Name: name, Surname: surname, Email: email, Age: age, Height: height, IsActive: isActive, CreationDate: creationDate}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	return p, nil
}
