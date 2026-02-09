package expenses

type Service struct {
	expenseList string
}

func NewService(list string) *Service {
	return &Service{
		expenseList: list,
	}
}
