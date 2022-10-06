package users

import (
	"fmt"
	"time"

	"github.com/MauriRossoMELI/backpack-bcgow6-mauricio-rosso/Documents/MAURI/BOOTCAMP/Go/backpack-bcgow6-mauricio-rosso/GO/go-web/PracticaDominios/pkg/store"
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
	Delete(id int) error
	UpdateSurnameAge(id int, surname string, age int) (User, error)
}

type repository struct {
	db store.Store
} //struct implementa los metodos de la interfaz

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]User, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, name string, surname string, email string, age int, height int, isActive bool, creationDate time.Time) (User, error) {

	//Lectura del json
	err := r.db.Read(&ps)
	if err != nil {
		return User{}, err
	}

	p := User{id, name, surname, email, age, height, isActive, creationDate}
	ps = append(ps, p)
	lastID = p.Id

	//Escritura en el json
	if err := r.db.Write(ps); err != nil {
		return User{}, err
	}

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

func (r *repository) UpdateSurnameAge(id int, surname string, age int) (User, error) {
	updated := false
	userUpdated := User{}
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Surname = surname
			ps[i].Age = age
			updated = true
			userUpdated = ps[i]
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	return userUpdated, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	for i := range ps {
		if ps[i].Id == id {
			ps = append(ps[0:i], ps[i+1:]...)
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Usuario %d no encontrado", id)
	}
	return nil
}
