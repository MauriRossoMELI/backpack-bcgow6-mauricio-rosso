package test

import (
	"encoding/json"

	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/internal/users"
)

type mockStorage struct {
	DataMock      []users.User
	ErrWrite      string
	ErrRead       string
	ReadWasCalled bool
}

func (s *mockStorage) Read(data interface{}) error {
	s.ReadWasCalled = true
	usersData := data.(*[]users.User)
	stubData := []users.User{
		// {
		// 	Id:           1,
		// 	Name:         "Mauri",
		// 	Surname:      "Rosso",
		// 	Email:        "mauri@mercadolibre.com",
		// 	Age:          1,
		// 	Height:       1293,
		// 	IsActive:     true,
		// 	CreationDate: "01-01-2012",
		// },
		{
			Id:           2,
			Name:         "Pablo",
			Surname:      "Testing",
			Email:        "pablo@mercadolibre.com",
			Age:          1,
			Height:       1293,
			IsActive:     true,
			CreationDate: "01-01-2012",
		},
	}
	*usersData = stubData
	return nil
}

func (s *mockStorage) Write(data interface{}) error {
	stubData := []users.User{
		{
			Id:           1,
			Name:         "Mauri",
			Surname:      "Rosso",
			Email:        "mauri@mercadolibre.com",
			Age:          1,
			Height:       1293,
			IsActive:     true,
			CreationDate: "01-01-2012",
		},
		{
			Id:           2,
			Name:         "Pablo",
			Surname:      "Testing",
			Email:        "pablo@mercadolibre.com",
			Age:          1,
			Height:       1293,
			IsActive:     true,
			CreationDate: "01-01-2012",
		},
	}
	stubDataB, errMarshal := json.Marshal(stubData)
	if errMarshal != nil {
		return errMarshal
	}
	errUnmarshal := json.Unmarshal(stubDataB, &data)
	if errUnmarshal != nil {
		return errUnmarshal
	}

	return nil
}
