package product

import (
	"context"
	"errors"
	"storageproject/internal/domain"
)

type Service interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	GetById(ctx context.Context, id int) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product, id int) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, id int) bool
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Exists(ctx context.Context, id int) bool {
	return s.repo.Exists(ctx, id)
}

func (s *service) GetByName(ctx context.Context, name string) (product domain.Product, err error) {
	product, err = s.repo.GetByName(ctx, name)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *service) GetById(ctx context.Context, id int) (product domain.Product, err error) {
	product, err = s.repo.GetById(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repo.GetAll(ctx)
	if err != nil {
		return []domain.Product{}, err
	}
	return products, err
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) Save(ctx context.Context, p domain.Product) (int, error) {
	if s.repo.Exists(ctx, p.Id) {
		return 0, errors.New("error: movie id already exists")
	}
	prod_id, err := s.repo.Save(ctx, p)
	if err != nil {
		return prod_id, err
	}

	p.Id = prod_id
	return p.Id, nil
}

func (s *service) Update(ctx context.Context, p domain.Product, id int) error {

	err := s.repo.Update(ctx, p, id)
	if err != nil {
		return err
	}
	_, errGetId := s.repo.GetById(ctx, id)
	if errGetId != nil {
		return errGetId
	}
	return nil
}
