package expenses

import (
	"encoding/json"
	"expense-tracker/internal/domain"
	"os"
	"path/filepath"
)

func (s *Service) SaveInformation(e []domain.Expense) error {
	dir := filepath.Dir(s.expenseList)
	if dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	file, err := os.Create(s.expenseList)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	return encoder.Encode(e)

}
