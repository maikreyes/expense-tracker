package expenses

import (
	"encoding/json"
	"errors"
	"expense-tracker/internal/domain"
	"io"
	"os"
)

func (s *Service) LoadInformation() ([]domain.Expense, error) {

	data, err := os.Open(s.expenseList)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []domain.Expense{}, nil
		}
		return nil, err
	}

	defer data.Close()

	var expenses []domain.Expense

	err = json.NewDecoder(data).Decode(&expenses)

	if err != nil {
		if errors.Is(err, io.EOF) {
			return []domain.Expense{}, nil
		}
		return nil, err
	}
	return expenses, nil

}
