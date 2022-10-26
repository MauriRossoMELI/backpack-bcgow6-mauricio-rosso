package users

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceUpdate(t *testing.T) {
	//arrange
	myStubStore := StubStore{}
	repo := NewRepository(&myStubStore)
	service := NewService(repo)
	dataEsperada := User{1, "Mauri", "Rosso", "mauri@mercadolibre.com", 1, 1293, true, "01-01-2012"}
	//expectedError := errors.New("usuario 1 no encontrado")
	//act
	resultado, err := service.Update(1, "Mauri", "Rosso", "mauri@mercadolibre.com", 1, 1293, true, "01-01-2012")
	//assert
	assert.Nil(t, err)
	assert.Equal(t, dataEsperada, resultado)
	assert.True(t, myStubStore.ReadWasCalled)
}

func TestServiceDelete(t *testing.T) {

	myStubStore := StubStore{}
	repo := NewRepository(&myStubStore)
	service := NewService(repo)

	// 1. Valida que user borrado efectivamente no exista. Se debe modificar el Read() para que funcione bien el GetAll()
	// arrange
	dataEsperada := []User{
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
	_ = service.Delete(1)
	users, _ := service.GetAll()
	//assert
	assert.Equal(t, dataEsperada, users)

	// 2. Obtiene el error correspondiente
	//arrange
	expectedError := errors.New("usuario 15 no encontrado")
	//act
	err := service.Delete(15)
	//assert
	assert.EqualError(t, err, expectedError.Error())
}

//AYUDA MEMORIA
// func TestServiceIntegrationStoreFail(t *testing.T) {
// 	// Arrange.
// 	expectedErr := errors.New("hello, i'm a little bug >:(")

// 	mockStorage := MockStorage{
// 		dataMock:   nil,
// 		errOnRead:  nil,
// 		errOnWrite: errors.New("hello, i'm a little bug >:("),
// 	}

// 	repository := NewRepository(&mockStorage)
// 	service := NewService(repository)

// 	// Act.
// 	userToCreate := User{
// 		Id:       1,
// 		Name:     "Mauri",
// 		Email:    "mauri@mercadolibre.cl",
// 		Age:      1,
// 		Height:   1293,
// 		IsActive: true,
// 	}

// 	result, err := service.Store(userToCreate)

// 	// Assert.
// 	assert.EqualError(t, err, expectedErr.Error())
// 	assert.Empty(t, result)
// }
