package get

type OrderGetter interface {
	GetOrderHandler(OrderUID string) (string, error)
}
