package service

type Service struct {
	OrderGetter
}

type OrderGetter interface {
	GetOrderByID(OrderUID string) (string, error)
}

func NewService() *Service {
	return &Service{}
}
