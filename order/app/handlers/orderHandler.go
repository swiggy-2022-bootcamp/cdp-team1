package handlers

import (
	"orderService/domain/service"
)

type OrderHandler struct {
	orderService service.OrderService
}
