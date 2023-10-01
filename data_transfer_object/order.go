package datatransferobject

import "time"

type NewOrderRequest struct {
	OrderAt      time.Time        `json:"ordered_at"`
	CustomerName string           `json:"customer_name"`
	Items        []NewItemRequest `json:"items"`
}

type UpdateResponse struct {
	Code int            `json:"code"`
	Data OrderWithItems `json:"data"`
}

type GetOrdersResponse struct {
	Data []OrderWithItems `json:"Data"`
}

type GetAnOrderResponse struct {
	Data OrderWithItems `json:"Data"`
}

type OrderWithItems struct {
	Id           int               `json:"id"`
	CustomerName string            `json:"customer_name"`
	OrderedAt    time.Time         `json:"ordered_at"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Items        []GetItemResponse `json:"Items"`
}

type GetItemResponse struct {
	Id           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ItemCode     string    `json:"itemcode"`
	Descriptions string    `json:"description"`
	Quantity     int       `json:"quantity"`
	OrderedId    int       `json:"orderid"`
}
