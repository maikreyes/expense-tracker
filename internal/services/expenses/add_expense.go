package expenses

import (
	"expense-tracker/internal/domain"
	"fmt"
	"time"
)

func (s *Service) AddExpense(description string, amount float64) error {

	data, err := s.LoadInformation()

	if err != nil {
		return err
	}

	id := 1

	if len(data) > 0 {
		id = data[len(data)-1].ID + 1
	}

	newExpense := domain.Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		Month:       time.Now().Month(),
	}

	data = append(data, newExpense)

	err = s.SaveInformation(data)

	if err != nil {
		return err
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", newExpense.ID)
	return nil
}
