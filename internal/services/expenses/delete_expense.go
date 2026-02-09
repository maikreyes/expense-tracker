package expenses

import (
	"expense-tracker/internal/domain"
	"fmt"
)

func (s *Service) DeleteExpense(id int) error {

	data, err := s.LoadInformation()

	if err != nil {
		return nil
	}

	var newData []domain.Expense

	for _, e := range data {
		if e.ID != id {
			newData = append(newData, e)
		}
	}

	err = s.SaveInformation(newData)

	if err != nil {
		fmt.Println("Error to try to eliminate expense with ID: ", id)
	}

	fmt.Println("Expense deleted successfully")

	return nil
}
