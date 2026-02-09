package expense

import (
	"fmt"
	"time"
)

func (h *Handler) GetSummary(month time.Month) error {

	total, err := h.Service.GetSummary(month)

	if err != nil {
		return err
	}

	if month != 0 {
		fmt.Printf("Total expenses for %s: %.2f\n", month, total)
	} else {
		fmt.Printf("Total expenses: %.2f\n", total)
	}

	return nil
}
