package mocks

import (
	"fmt"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/domain"
)

// NOT USED FOR THIS EXERCISE
type MockService struct {
	DataMock []domain.Product
	Error    string
}

func (m *MockService) GetAllBySeller() ([]domain.Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil

}
