package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mmfshirokan/accountService/internal/model"
	"github.com/mmfshirokan/accountService/internal/repository"
)

type service struct {
	repo repository.Interface
}

func New(repo repository.Interface) Interface {
	return &service{
		repo: repo,
	}
}

type Interface interface {
	Create(id uuid.UUID) error
	PayIn(ctx context.Context, payIn model.Balance) error
	PayOut(ctx context.Context, payOut model.Balance) error
	Get(ctx context.Context, id uuid.UUID) (model.Balance, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func (s *service) Create(id uuid.UUID) error {
	return s.repo.Create(id)
}

func (s *service) PayIn(ctx context.Context, payIn model.Balance) error {
	return s.repo.PayIn(ctx, payIn)
}

func (s *service) PayOut(ctx context.Context, payOut model.Balance) error {
	return s.repo.PayOut(ctx, payOut)
}

func (s *service) Get(ctx context.Context, id uuid.UUID) (model.Balance, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
