package model

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint
	Name      string
	Price     int64
	SKU       string `gorm:"uniqueIndex"` // Stock Keeping Unit, unique identifier
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GenerateProducts() []Product {
	products := make([]Product, 10)
	now := time.Now()
	for i := 0; i < 10; i++ {
		// 生成唯一SKU：时间戳+随机数+序号
		sku := fmt.Sprintf("SKU%s%03d%02d", now.Format("20060102150405"), rand.Intn(999), i+1)
		products[i] = Product{
			Name:      fmt.Sprintf("商品%d", i+1),
			Price:     int64((i + 1) * 100), // 100, 200, 300...
			SKU:       sku,
			CreatedAt: now,
			UpdatedAt: now,
		}
	}
	return products
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.SKU == "" { // 如果SKU为空，则生成一个默认的SKU
		p.SKU = fmt.Sprintf("SKU%s%03d", time.Now().Format("20060102150405"), rand.Intn(999))
	}

	if p.Price < 0 {
		return fmt.Errorf("价格不能为负")
	}

	return nil
}
