package ports

import (
	"expense-tracker/internal/domain"
	"time"
)

type ExpenseService interface {
	AddExpense(description string, amount float64) error
	ListExpenses() ([]domain.Expense, error)
	DeleteExpense(id int) error
	GetSummary(month time.Month) (float64, error)
}
