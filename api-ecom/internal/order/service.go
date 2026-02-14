package order

import (
	repo "api-ecom/internal/adapters/psql/sqlc"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder placeOrderPrams) (repo.Order, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) *svc {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder placeOrderPrams) (repo.Order, error) {
	if tempOrder.UserId == 0 {
		return repo.Order{}, errors.New("invalid user_id: " + string(tempOrder.UserId))
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, errors.New("empty items")
	}

	tx, err := s.db.Begin(ctx) // create transaction
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx) // for transaction queries use this

	order, err := s.repo.CrateOrder(ctx, tempOrder.UserId)
	if err != nil {
		return repo.Order{}, err
	}

	for _, item := range tempOrder.Items {
		p, err := qtx.GetProductById(ctx, item.ProductId)
		if err != nil {
			return repo.Order{}, errors.New("product not found")
		}

		if p.Quantity < int32(item.Quantity) {
			return repo.Order{}, errors.New(string(item.ProductId) + " product has not enough stock")
		}

		_, err = qtx.CrateOrderItems(ctx, repo.CrateOrderItemsParams{
			OrderID:   order.ID,
			ProductID: item.ProductId,
			Price:     p.Price,
			Quantity:  int32(item.Quantity),
		})
		if err != nil {
			return repo.Order{}, err
		}
	}

	tx.Commit(ctx) // finish transaction and commit

	return order, nil
}
