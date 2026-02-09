package expenses

import (
	"expense-tracker/internal/domain"
)

func (s *Service) ListExpenses() ([]domain.Expense, error) {

	data, err := s.LoadInformation()

	if err != nil {
		return nil, err
	}

	return data, nil
}
