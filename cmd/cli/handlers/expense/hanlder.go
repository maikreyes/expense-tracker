package expense

import (
	"expense-tracker/internal/ports"
)

type Handler struct {
	Service ports.ExpenseService
}

func NewHandler(s ports.ExpenseService) *Handler {
	return &Handler{
		Service: s,
	}
}
