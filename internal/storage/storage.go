package storage

import (
	"context"
	"sync"

	"github.com/google/uuid"

	"calculate-distributed/internal/api"
	"calculate-distributed/internal/ownErrors"
)

var _ Storage = (*storage)(nil)

type storage struct {
	expressions map[*uuid.UUID]api.Expression
	tasks       map[*uuid.UUID]api.GetTaskResponse
	mu          sync.Mutex
}

type Storage interface {
	AddExpression(ctx context.Context, expression api.Expression) error
	GetExpression(ctx context.Context, id *uuid.UUID) (api.Expression, error)
	UpdateExpression(ctx context.Context, expression api.Expression) error
	AddTask(ctx context.Context, task api.GetTaskResponse) error
	GetTask(ctx context.Context, id *uuid.UUID) (api.GetTaskResponse, error)
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

func (s *storage) GetExpression(ctx context.Context, id *uuid.UUID) (api.Expression, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) UpdateExpression(ctx context.Context, expression api.Expression) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) AddTask(ctx context.Context, task api.GetTaskResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) GetTask(ctx context.Context, id *uuid.UUID) (api.GetTaskResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) UpdateTask(ctx context.Context, task api.GetTaskResponse) error {
	//TODO implement me
	panic("implement me")
}
