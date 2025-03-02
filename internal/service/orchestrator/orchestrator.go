package orchestrator

import (
	"context"

	"github.com/google/uuid"

	"calculate-distributed/internal/api"
	"calculate-distributed/internal/storage"
)

type orchestrator struct {
	storage storage.Storage
}

type Orchestrator interface {
	AddExpression(ctx context.Context, exprString string) (uuid.UUID, error)
}

func New(s storage.Storage) Orchestrator {
	return &orchestrator{
		storage: s,
	}
}

func (o *orchestrator) AddExpression(ctx context.Context, exprString string) (uuid.UUID, error) {
	exprUUID := uuid.New()
	status := api.Accepted

	expr := api.Expression{
		Id:     &exprUUID,
		Result: nil,
		Status: &status,
	}

	if err := o.storage.AddExpression(ctx, expr); err != nil {
		return uuid.Nil, err
	}

	return exprUUID, nil
}
