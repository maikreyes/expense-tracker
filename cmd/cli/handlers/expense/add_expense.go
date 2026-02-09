package expense

func (h *Handler) AddExpense(description string, amount float64) error {

	err := h.Service.AddExpense(description, amount)

	if err != nil {
		return err
	}

	return nil
}
