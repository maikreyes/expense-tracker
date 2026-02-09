package expenses

import "time"

func (s *Service) GetSummary(month time.Month) (float64, error) {

	data, err := s.LoadInformation()

	if err != nil {
		return 0.0, err
	}

	var total float64

	if month != 0 {
		for _, e := range data {
			if e.Month == month {
				total += e.Amount
			}
		}

		return total, nil
	}

	for _, e := range data {
		total += e.Amount
	}

	return total, nil
}
