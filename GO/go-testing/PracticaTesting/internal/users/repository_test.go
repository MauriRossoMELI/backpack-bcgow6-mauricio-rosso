package users

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	ReadWasCalled bool
}

func (s *StubStore) Read(data interface{}) error {
	s.ReadWasCalled = true

	users := data.(*[]User)
	stubData := []User{
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
	*users = stubData
	return nil
}

func (s StubStore) Write(data interface{}) error {
	stubData := []User{
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
	errUnmarshal := json.Unmarshal(stubDataB, data)
	if errUnmarshal != nil {
		return errUnmarshal
	}

	return nil
}

func TestGetAll(t *testing.T) {
	//arrange
	myStubStore := StubStore{}
	repo := NewRepository(&myStubStore)
	dataEsperada := []User{
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
	//act
	resultado, _ := repo.GetAll()
	//assert
	assert.Equal(t, dataEsperada, resultado)
}
