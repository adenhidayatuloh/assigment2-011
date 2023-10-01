package orderrepository

import "github.com/adenhidayatuloh/assigment2-011/entity"

type OrderItem struct {
	Order entity.Order
	Item  entity.Item
}

type OrderItemMapped struct {
	Order entity.Order
	Items []entity.Item
}

func (oim *OrderItemMapped) HandleMappingOrderWithItems(orderitem []OrderItem) []OrderItemMapped {
	ordersItemMapped := []OrderItemMapped{}

	for _, eachOrderItem := range orderitem {

		isOrderExist := false

		for i := range ordersItemMapped {

			if eachOrderItem.Order.OrderId == ordersItemMapped[i].Order.OrderId {
				isOrderExist = true
				ordersItemMapped[i].Items = append(ordersItemMapped[i].Items, eachOrderItem.Item)
				break
			}
		}

		if !isOrderExist {
			orderItemMapped := OrderItemMapped{
				Order: eachOrderItem.Order,
			}
			orderItemMapped.Items = append(orderItemMapped.Items, eachOrderItem.Item)

			ordersItemMapped = append(ordersItemMapped, orderItemMapped)
		}
	}
	return ordersItemMapped
}
