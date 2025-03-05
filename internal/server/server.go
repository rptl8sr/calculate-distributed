package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	openapitypes "github.com/oapi-codegen/runtime/types"

	"calculate-distributed/internal/api"
	"calculate-distributed/internal/ownErrors"
	"calculate-distributed/internal/service/orchestrator"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	orchestrator service.Service
}

func (s Server) PostApiV1Calculate(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()
	if err != nil {
		respondError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}
	defer func() { _ = body.Close() }()

	var expr api.ExpressionRequest
	if err = json.NewDecoder(body).Decode(&expr); err != nil {
		respondError(w, http.StatusBadRequest, "Failed to parse request body")
		return
	}

	if expr.Expression == nil {
		respondError(w, http.StatusBadRequest, "Expression is null")
		return
	}

	if *expr.Expression == "" {
		respondError(w, http.StatusBadRequest, "Expression is empty")
		return
	}

	exprUUID, err := s.orchestrator.AddExpression(r.Context(), *expr.Expression)
	if err != nil {
		switch {
		case errors.Is(err, ownErrors.ErrExpressionExists):
			respondError(w, http.StatusBadRequest, err.Error())
			return
		case errors.Is(err, ownErrors.ErrIDNotFound):

		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, api.ExpressionAccepted{Id: &exprUUID})
}

func (s Server) GetApiV1Expressions(w http.ResponseWriter, r *http.Request) {
	exprList, err := s.orchestrator.ExpressionsList(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, api.ExpressionsList{Expressions: &exprList})
}

func (s Server) GetApiV1ExpressionsId(w http.ResponseWriter, r *http.Request, id openapitypes.UUID) {
	if id == uuid.Nil {
		respondError(w, http.StatusBadRequest, "ID is null")
		return
	}

	expr, err := s.orchestrator.Expression(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, ownErrors.ErrIDNotFound):
			respondError(w, http.StatusNotFound, err.Error())
			return
		default:
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondJSON(w, http.StatusOK, expr)
}

func (s Server) GetInternalTask(w http.ResponseWriter, r *http.Request) {
	task, err := s.orchestrator.Task(r.Context())
	if err != nil {

	}

	respondJSON(w, http.StatusOK, task)
}

func (s Server) PostInternalTask(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func New(o service.Service) *Server {
	return &Server{
		orchestrator: o,
	}
}
