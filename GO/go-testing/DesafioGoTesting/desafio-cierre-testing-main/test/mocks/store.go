package mocks

import (
	"fmt"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/domain"
)

// NOT USED FOR THIS EXERCISE
type MockStorage struct {
	DataMock []domain.Product
	ErrWrite string
	ErrRead  string
}

func (m *MockStorage) Read(data interface{}) error {
	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}
	a := data.(*[]domain.Product)
	*a = m.DataMock
	return nil
}
