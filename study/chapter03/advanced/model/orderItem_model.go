package model

import "time"

type OrderItem struct {
	ID        uint
	OrderID   uint    // Foreign key to order
	ProductID uint    // Foreign key to product
	Product   Product // Belongs To: OrderItem belongs to one product
	Quantity  int
	UnitPrice int64
	CreatedAt time.Time
}

func GenerateOrderItems(orders []Order, products []Product) []OrderItem {
	orderItems := make([]OrderItem, len(orders))
	now := time.Now()
	for i, order := range orders {
		product := products[i%10] // 循环关联商品
		orderItems[i] = OrderItem{
			OrderID:   order.ID,   // 关联真实Order.ID
			ProductID: product.ID, // 关联真实Product.ID
			Product:   product,
			Quantity:  i + 1, // 1,2,3...
			UnitPrice: product.Price,
			CreatedAt: now,
		}
	}
	return orderItems
}
