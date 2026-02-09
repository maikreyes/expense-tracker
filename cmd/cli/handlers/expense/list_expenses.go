package expense

import "fmt"

func (h *Handler) ListExpenses() error {

	data, err := h.Service.ListExpenses()

	if err != nil {
		return err
	}

	fmt.Println("Id description amount moth ")

	for _, e := range data {
		fmt.Printf("%v, %v, %v, %v\n", e.ID, e.Description, e.Amount, e.Month.String())
	}

	return nil
}
