package expense

func (h *Handler) DeleteExpense(id int) error {

	err := h.Service.DeleteExpense(id)

	if err != nil {
		return err
	}

	return nil
}
