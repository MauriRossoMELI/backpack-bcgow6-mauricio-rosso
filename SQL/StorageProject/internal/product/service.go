package product

import (
	"context"
	"storageproject/internal/domain"
)

type Service interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
	// Save(ctx context.Context, b domain.Movie) (domain.Movie, error)
	// Update(ctx context.Context, b domain.Movie, id int) (domain.Movie, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetByName(ctx context.Context, name string) (product domain.Product, err error) {
	product, err = s.repo.GetByName(ctx, name)
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

/*
func (s *service) Save(ctx context.Context, m domain.Movie) (domain.Movie, error) {
	if s.repo.Exists(ctx, m.ID) {
		return domain.Movie{}, errors.New("error: movie id already exists")
	}
	movie_id, err := s.repo.Save(ctx, m)
	if err != nil {
		return domain.Movie{}, err
	}

	m.ID = int(movie_id)
	return m, nil
}

func (s *service) Update(ctx context.Context, b domain.Movie, id int) (domain.Movie, error) {

	err := s.repo.Update(ctx, b, id)
	if err != nil {
		return domain.Movie{}, err
	}
	updated, err := s.repo.Get(ctx, id)
	if err != nil {
		return b, err
	}
	return updated, nil
}
*/
