package service

import (
	datatransferobject "assigment2_aden/data_transfer_object"
	"assigment2_aden/entity"
	orderrepository "assigment2_aden/repository/order_repository"
	"errors"
)

type orderService struct {
	OrderRepo orderrepository.Repository
}

type OrderService interface {
	CreateOrder(newOrderRequest datatransferobject.NewOrderRequest) error
	GetOrders() (*datatransferobject.GetOrdersResponse, error)
	GetOrderByID(OrdersID int) (*datatransferobject.GetAnOrderResponse, error)
	UpdateOrder(orderID int, newOrderRequest datatransferobject.NewOrderRequest) error
	DeleteOrder(orderID int) error
}

func NewOrderService(OrderRepo orderrepository.Repository) OrderService {
	return &orderService{
		OrderRepo: OrderRepo,
	}
}

func (os *orderService) CreateOrder(newOrderRequest datatransferobject.NewOrderRequest) error {

	orderPayload := entity.Order{
		OrderedAt:    newOrderRequest.OrderAt,
		CustomerName: newOrderRequest.CustomerName,
	}

	itemPayload := []entity.Item{}

	for _, eachItem := range newOrderRequest.Items {
		item := entity.Item{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
		}
		itemPayload = append(itemPayload, item)
	}
	err := os.OrderRepo.CreateOrder(orderPayload, itemPayload)

	if err != nil {
		return err
	}
	return nil
}

func (os *orderService) GetOrders() (*datatransferobject.GetOrdersResponse, error) {
	orders, err := os.OrderRepo.ReadOrders()

	if err != nil {
		return nil, err
	}

	orderResult := []datatransferobject.OrderWithItems{}

	for _, eachOrder := range orders {
		itemResult := []datatransferobject.GetItemResponse{}
		order := datatransferobject.OrderWithItems{

			CustomerName: eachOrder.Order.CustomerName,
			Id:           eachOrder.Order.OrderId,
			OrderedAt:    eachOrder.Order.OrderedAt,
			CreatedAt:    eachOrder.Order.CreatedAt,
			UpdatedAt:    eachOrder.Order.UpdatedAt,
		}

		for _, eachItems := range eachOrder.Items {

			item := datatransferobject.GetItemResponse{
				Id:           eachItems.ItemId,
				CreatedAt:    eachItems.CreatedAt,
				UpdatedAt:    eachItems.UpdatedAt,
				ItemCode:     eachItems.ItemCode,
				Descriptions: eachItems.Description,
				Quantity:     eachItems.Quantity,
				OrderedId:    eachItems.OrderId,
			}

			itemResult = append(itemResult, item)

		}
		order.Items = itemResult
		orderResult = append(orderResult, order)

	}

	getOrdersResponse := datatransferobject.GetOrdersResponse{
		Data: orderResult,
	}

	return &getOrdersResponse, nil

}

func (os *orderService) GetOrderByID(OrdersID int) (*datatransferobject.GetAnOrderResponse, error) {

	getOrdersResponse, err := os.GetOrders()

	getAnOrder := datatransferobject.GetAnOrderResponse{}

	isFound := false

	if err != nil {
		return nil, err
	}

	for i := range getOrdersResponse.Data {

		if getOrdersResponse.Data[i].Id == OrdersID {
			getAnOrder.Data = getOrdersResponse.Data[i]
			isFound = true
			break
		}

	}

	if !isFound {
		return nil, errors.New("data not found")
	}
	return &getAnOrder, nil
}

func (os *orderService) UpdateOrder(orderID int, newOrderRequest datatransferobject.NewOrderRequest) error {

	//updateOrderQuery = `update "orders" set "ordered_at" = $1 , "customer_name" = $2 where "order_id" = $3 RETURNING "order_id"`
	//updateItemQuery  = `update "items"  set "description"  = $1 , "quantity"  = $2 where "item_code"  = $3 and "order_id" = $4;`

	//err = tx.QueryRow(CreateOrderQuery, orderPayload.OrderedAt, orderPayload.CustomerName).Scan(&orderId)
	//tx.Exec(CreateItemQuery, eachItem.ItemCode, eachItem.Description, eachItem.Quantity, orderId)

	orderPayload := entity.Order{
		OrderedAt:    newOrderRequest.OrderAt,
		CustomerName: newOrderRequest.CustomerName,
		OrderId:      orderID,
	}

	itemPayload := []entity.Item{}

	for _, eachItem := range newOrderRequest.Items {
		item := entity.Item{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
			OrderId:     orderID,
		}
		itemPayload = append(itemPayload, item)
	}
	err := os.OrderRepo.UpdateOrder(orderPayload, itemPayload)

	if err != nil {
		return err
	}
	return nil
}

func (os *orderService) DeleteOrder(orderID int) error {

	OrderDelete := entity.Order{
		OrderId: orderID,
	}
	err := os.OrderRepo.DeleteOrder(OrderDelete)

	if err != nil {
		return err
	}

	return nil

}
