package users

import "time"

type Service interface {
	GetAll() ([]User, error)
	Store(name, surname, email string, age, height int, isActive bool, creationDate time.Time) (User, error)
	Update(id int, name, surname, email string, age, height int, isActive bool, creationDate time.Time) (User, error)
	Delete(id int) error
	UpdateSurnameAge(id int, surname string, age int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]User, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(name string, surname string, email string, age int, height int, isActive bool, creationDate time.Time) (User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return User{}, err
	}

	lastID++

	user, err := s.repository.Store(lastID, name, surname, email, age, height, isActive, creationDate)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) Update(id int, name string, surname string, email string, age int, height int, isActive bool, creationDate time.Time) (User, error) {

	return s.repository.Update(id, name, surname, email, age, height, isActive, creationDate)
}

func (s *service) UpdateSurnameAge(id int, surname string, age int) (User, error) {

	return s.repository.UpdateSurnameAge(id, surname, age)
}

func (s *service) Delete(id int) error {

	return s.repository.Delete(id)
}
