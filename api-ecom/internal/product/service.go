package product

import (
	"context"

	repo "api-ecom/internal/adapters/psql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	CreateProduct(ctx context.Context, arg repo.CreateProductParams) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (svc *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return svc.repo.ListProducts(ctx)
}

func (svc *svc) CreateProduct(ctx context.Context, arg repo.CreateProductParams) (repo.Product, error) {
	return svc.repo.CreateProduct(ctx, arg)
}
