package products

import (
	"errors"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-testing/test/mocks"
	"github.com/stretchr/testify/assert"
)

type mockedRepository struct{}

func NewMockedRepository() Repository {
	return &mockedRepository{}
}

func (mr *mockedRepository) GetAllBySeller(SellerID string) ([]domain.Product, error) {
	return nil, errors.New("Error!!!")
}

func TestServiceRepoGetAllOk(t *testing.T) {
	// arrange
	database := []domain.Product{
		{
			ID:          "mock1",
			SellerID:    "asd",
			Description: "generic product",
			Price:       123.55,
		}, {
			ID:          "mock2",
			SellerID:    "asd",
			Description: "standar product",
			Price:       333.33,
		}}
	mockStorage := mocks.MockStorage{
		DataMock: database,
		ErrWrite: "",
		ErrRead:  "",
	}
	// act
	repo := NewRepository()
	service := NewService(repo)
	result, err := service.GetAllBySeller("asd")
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.DataMock, result)
}

func TestServiceRepoGetAllFail(t *testing.T) {
	//arrange
	//expectedErr := errors.New("forced error in repository")
	var expected []domain.Product
	// act
	repo := NewRepository()
	service := NewService(repo)
	result, _ := service.GetAllBySeller("xd")
	//assert
	assert.Equal(t, expected, result)
	//assert.EqualError(t, err, expectedErr.Error()) //It needed to force an error in repo to test it
	//assert.True(t, mockProductRepository.wasInvoked)
}
