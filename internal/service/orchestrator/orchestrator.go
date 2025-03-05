package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	openapitypes "github.com/oapi-codegen/runtime/types"

	"calculate-distributed/internal/api"
	"calculate-distributed/internal/storage"
)

type service struct {
	storage   storage.Storage
	timeouts  map[api.Operation]time.Duration
	exprQueue chan openapitypes.UUID
	taskQueue chan openapitypes.UUID
}

type Service interface {
	AddExpression(ctx context.Context, exprString string) (uuid.UUID, error)
	ExpressionsList(ctx context.Context) ([]api.Expression, error)
	Expression(ctx context.Context, id openapitypes.UUID) (api.Expression, error)
	Task(ctx context.Context) (*api.GetTaskResponse, error)
}

func New(s storage.Storage, t map[api.Operation]time.Duration) Service {
	return &service{
		storage:   s,
		timeouts:  t,
		exprQueue: make(chan openapitypes.UUID, 1000),
		taskQueue: make(chan openapitypes.UUID, 1000),
	}
}

func (s *service) AddExpression(ctx context.Context, exprString string) (uuid.UUID, error) {
	exprUUID := uuid.New()
	status := api.Accepted

	expr := api.Expression{
		Id:     &exprUUID,
		Result: nil,
		Status: &status,
	}

	if err := s.storage.AddExpression(ctx, expr); err != nil {
		return uuid.Nil, err
	}
	// TODO: add expression workflow

	s.exprQueue <- exprUUID

	return exprUUID, nil
}

func (s *service) ExpressionsList(ctx context.Context) ([]api.Expression, error) {
	return s.storage.ExpressionsList(ctx)
}

func (s *service) Expression(ctx context.Context, id openapitypes.UUID) (api.Expression, error) {
	return s.storage.Expression(ctx, id)
}

func (s *service) Task(ctx context.Context) (*api.GetTaskResponse, error) {
	//task, err := s.storage.Task(ctx)
	return nil, nil
}
