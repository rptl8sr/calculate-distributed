package storage

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"

	"calculate-distributed/internal/api"
	"calculate-distributed/internal/ownErrors"
)

var _ Storage = (*storage)(nil)

type storage struct {
	operationTime map[api.Operation]time.Duration
	expressions   sync.Map
	tasks         sync.Map
	//expressions   map[*uuid.UUID]api.Expression
	//tasks         map[*uuid.UUID]api.GetTaskResponse
}

type Storage interface {
	AddExpression(ctx context.Context, expression api.Expression) error
	Expression(ctx context.Context, id uuid.UUID) (api.Expression, error)
	ExpressionsList(ctx context.Context) ([]api.Expression, error)
	UpdateExpression(ctx context.Context, expression api.Expression) error
	AddTask(ctx context.Context, task api.GetTaskResponse) error
	Task(ctx context.Context) (api.GetTaskResponse, error)
	UpdateTask(ctx context.Context, task api.GetTaskResponse) error
}

func New() Storage {
	return &storage{
		expressions: make(map[*uuid.UUID]api.Expression),
		tasks:       make(map[*uuid.UUID]api.GetTaskResponse),
	}
}

func (s *storage) AddExpression(_ context.Context, e api.Expression) error {
	if e.Id == nil {
		return ownErrors.ErrIDNotFound
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.expressions[e.Id]; ok {
		return ownErrors.ErrExpressionExists
	}

	s.expressions[e.Id] = e
	return nil
}

func (s *storage) Expression(ctx context.Context, id uuid.UUID) (api.Expression, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	expr, ok := s.expressions[&id]
	if !ok {
		return api.Expression{}, ownErrors.ErrIDNotFound
	}

	return expr, nil
}

func (s *storage) ExpressionsList(_ context.Context) ([]api.Expression, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	exprs := make([]api.Expression, 0, len(s.expressions))
	for _, expr := range s.expressions {
		exprs = append(exprs, expr)
	}

	return exprs, nil
}

func (s *storage) UpdateExpression(ctx context.Context, expression api.Expression) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) AddTask(ctx context.Context, task api.GetTaskResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) Task(ctx context.Context) (api.GetTaskResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) UpdateTask(ctx context.Context, task api.GetTaskResponse) error {
	//TODO implement me
	panic("implement me")
}
