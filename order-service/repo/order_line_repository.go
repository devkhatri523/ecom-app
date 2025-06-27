package repo

import "order-service/domain"

type OrderLineRepository interface {
	SaveOrderLine(orderLine domain.OrderLine) error
}
