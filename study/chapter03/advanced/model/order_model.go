package model

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID         uint
	OrderNo    string      `gorm:"uniqueIndex"` // Unique order number
	UserID     uint        // Foreign key to user
	Items      []OrderItem // Has Many: One order has many items
	TotalPrice int64
	Status     string
	CreatedAt  time.Time
}

func GenerateOrders(users []User) []Order {
	orders := make([]Order, 10)
	now := time.Now()
	for i := 0; i < 10; i++ {
		// 生成唯一订单号：纳秒时间戳+随机数+序号
		orderNo := fmt.Sprintf("ORD%d%05d%02d", now.UnixNano(), rand.Intn(99999), i+1)
		orders[i] = Order{
			OrderNo:    orderNo,
			UserID:     users[i%10].ID,                 // 关联真实User.ID
			TotalPrice: int64((i + 1) * (i + 1) * 100), // 100, 400, 900...
			Status:     []string{"pending", "paid", "shipped", "delivered"}[i%4],
			CreatedAt:  now,
		}
	}
	return orders
}
