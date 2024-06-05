package get

type OrderGetter interface {
	GetOrder(request_id string) (string, error)
}
