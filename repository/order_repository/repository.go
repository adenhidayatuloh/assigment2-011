package orderrepository

import "github.com/adenhidayatuloh/assigment2-011/entity"

type Repository interface {
	CreateOrder(orderPayload entity.Order, itemPayload []entity.Item) error
	ReadOrders() ([]OrderItemMapped, error)
	UpdateOrder(orderPayload entity.Order, itemPayload []entity.Item) error
	DeleteOrder(orderDelete entity.Order) error
}
