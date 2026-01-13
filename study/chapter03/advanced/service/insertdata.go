package service

import (
	"fmt"

	"gorm.io/gorm"

	"local/go_study/study/chapter03/advanced/model"
)

// type GeneratedData struct {
// 	Products   []model.Product
// 	Roles      []model.Role
// 	Profiles   []model.Profile
// 	OrderItems []model.OrderItem
// 	Orders     []model.Order
// 	Users      []model.User
// }

// func GenerateData() GeneratedData {
// 	// 生成基础数据
// 	products := model.GenerateProducts()
// 	roles := model.GenerateRoles()

// 	tempUsers := make([]model.User, 10)
// 	for i := 0; i < 10; i++ {
// 		tempUsers[i].ID = uint(i + 1)
// 	}

// 	profiles := model.GenerateProfiles(tempUsers)

// 	tempOrders := make([]model.Order, 10)
// 	for i := 0; i < 10; i++ {
// 		tempOrders[i].ID = uint(i + 1)
// 		tempOrders[i].UserID = uint((i % 10) + 1)
// 	}

// 	// 生成关联数据
// 	orderItems := model.GenerateOrderItems(tempOrders, products)
// 	orders := model.GenerateOrders(tempUsers, orderItems)
// 	users := model.GenerateUsers(profiles, orders, roles)

// 	// 返回所有生成的数据
// 	return GeneratedData{
// 		Products:   products,
// 		Roles:      roles,
// 		Profiles:   profiles,
// 		OrderItems: orderItems,
// 		Orders:     orders,
// 		Users:      users,
// 	}
// }

func InsertData(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ========== 步骤1：插入无依赖的基础数据 ==========
	// 1.1 插入角色
	roles := model.GenerateRoles()
	if err := tx.Create(&roles).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("插入角色失败: %v", err)
	}

	// 1.2 插入商品
	products := model.GenerateProducts()
	if err := tx.Create(&products).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("插入商品失败: %v", err)
	}

	// ========== 步骤2：插入用户（纯数据，无关联） ==========
	users := model.GenerateUsers()
	if err := tx.Create(&users).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("插入用户失败: %v", err)
	}

	// ========== 步骤3：基于用户真实ID插入关联数据 ==========
	// 3.1 插入用户档案
	profiles := model.GenerateProfiles(users)
	if err := tx.Create(&profiles).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("插入用户档案失败: %v", err)
	}

	// 3.2 插入订单
	orders := model.GenerateOrders(users)
	if err := tx.Create(&orders).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("插入订单失败: %v", err)
	}

	// 3.3 插入订单项
	orderItems := model.GenerateOrderItems(orders, products)
	if err := tx.Create(&orderItems).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("插入订单项失败: %v", err)
	}

	// ========== 步骤4：关联用户和角色（多对多） ==========
	for i, user := range users {
		// 为每个用户分配2个角色
		role1 := roles[i%len(roles)]
		role2 := roles[(i+1)%len(roles)]
		if err := tx.Model(&user).Association("Roles").Replace(role1, role2); err != nil {
			tx.Rollback()
			return fmt.Errorf("关联用户%d和角色失败: %v", i+1, err)
		}
	}

	// ========== 步骤5：提交事务 ==========
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交插入事务失败: %v", err)
	}

	fmt.Println("所有数据插入成功！")
	return nil
}
