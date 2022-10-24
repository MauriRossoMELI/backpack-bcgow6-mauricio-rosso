package products

import "github.com/bootcamp-go/desafio-cierre-testing/internal/domain"

type Repository interface {
	GetAllBySeller(sellerID string) ([]domain.Product, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	var prodList []domain.Product
	var pordListFiltered []domain.Product

	prodList = append(prodList, domain.Product{
		ID:          "mock1",
		SellerID:    "asd",
		Description: "generic product",
		Price:       123.55,
	})
	prodList = append(prodList, domain.Product{
		ID:          "mock2",
		SellerID:    "asd",
		Description: "standar product",
		Price:       333.33,
	})
	for _, v := range prodList {
		if v.SellerID == sellerID {
			pordListFiltered = append(pordListFiltered, v)
		}
	}

	return pordListFiltered, nil
}
