package orderpostgres

import (
	"assigment2_aden/entity"
	orderrepository "assigment2_aden/repository/order_repository"
	"database/sql"
)

type orderPG struct {
	db *sql.DB
}

const (
	CreateOrderQuery = `insert into "orders"("ordered_at","customer_name") values ($1,$2) RETURNING "order_id"`

	CreateItemQuery = `insert into "items"("item_code","description","quantity","order_id")values($1,$2,$3,$4)`

	getOrdersWithItemsQuery = `select "o"."order_id","o"."customer_name","o"."ordered_at","o"."created_at","o"."updated_at",
	"i"."item_id","i"."item_code","i"."quantity","i"."description","i"."order_id","i"."created_at","i"."updated_at"
	from "orders" as "o"
	left join "items" as "i" on "o"."order_id" = "i"."order_id"
	order by "o"."order_id" asc`

	updateOrderQuery = `update "orders" set "ordered_at" = $1 , "customer_name" = $2,"updated_at" = NOW() where "order_id" = $3 `
	updateItemQuery  = `update "items"  set "description"  = $1 , "quantity"  = $2, "updated_at" = NOW()  where "item_code"  = $3 and "order_id" = $4;`
	deleteOrder      = `DELETE FROM "orders"  WHERE "order_id" = $1;`
)

func NewOrderPG(db *sql.DB) orderrepository.Repository {
	return &orderPG{db: db}
}

func (orderPG *orderPG) CreateOrder(orderPayload entity.Order, itemPayload []entity.Item) error {

	tx, err := orderPG.db.Begin()

	if err != nil {
		return err
	}

	var orderId int

	err = tx.QueryRow(CreateOrderQuery, orderPayload.OrderedAt, orderPayload.CustomerName).Scan(&orderId)

	if err != nil {
		tx.Rollback()
		return err

	}

	for _, eachItem := range itemPayload {
		_, err := tx.Exec(CreateItemQuery, eachItem.ItemCode, eachItem.Description, eachItem.Quantity, orderId)

		if err != nil {

			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (orderPG *orderPG) ReadOrders() ([]orderrepository.OrderItemMapped, error) {
	rows, err := orderPG.db.Query(getOrdersWithItemsQuery)

	if err != nil {
		return nil, err
	}

	orderItems := []orderrepository.OrderItem{}

	for rows.Next() {
		var orderItem orderrepository.OrderItem

		err = rows.Scan(
			&orderItem.Order.OrderId, &orderItem.Order.CustomerName, &orderItem.Order.OrderedAt, &orderItem.Order.CreatedAt, &orderItem.Order.UpdatedAt,
			&orderItem.Item.ItemId, &orderItem.Item.ItemCode, &orderItem.Item.Quantity, &orderItem.Item.Description, &orderItem.Item.OrderId, &orderItem.Item.CreatedAt, &orderItem.Item.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)

	}

	var result orderrepository.OrderItemMapped

	return result.HandleMappingOrderWithItems(orderItems), nil
}

func (orderPG *orderPG) UpdateOrder(orderPayload entity.Order, itemPayload []entity.Item) error {

	tx, err := orderPG.db.Begin()

	if err != nil {
		return err
	}
	//updateOrderQuery = `update "orders" set "ordered_at" = $1 , "customer_name" = $2,"updated_at" = $3 where "order_id" = $4 `
	_, err = tx.Exec(updateOrderQuery, orderPayload.OrderedAt, orderPayload.CustomerName, orderPayload.OrderId)

	if err != nil {
		tx.Rollback()
		return err

	}

	//updateItemQuery  = `update "items"  set "description"  = $1 , "quantity"  = $2, "updated_at" = $3 , where "item_code"  = $4 and "order_id" = $5;`

	for _, eachItem := range itemPayload {
		_, err := tx.Exec(updateItemQuery, eachItem.Description, eachItem.Quantity, eachItem.ItemCode, eachItem.OrderId)

		if err != nil {

			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (orderPG *orderPG) DeleteOrder(orderDelete entity.Order) error {

	tx, err := orderPG.db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(deleteOrder, orderDelete.OrderId)

	if err != nil {
		tx.Rollback()
		return err

	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil

}
