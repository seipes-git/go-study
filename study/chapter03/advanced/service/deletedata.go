package service

import (
	"fmt"

	"gorm.io/gorm"

	"local/go_study/study/chapter03/advanced/model"
)

func DeleteData(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Exec("DELETE FROM gva_user_roles").Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("清空用户角色中间表失败: %v", err)
	}

	if err := tx.Model(&model.OrderItem{}).Unscoped().Where("1=1").Delete(&model.OrderItem{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("清空订单项失败: %v", err)
	}

	if err := tx.Model(&model.Order{}).Unscoped().Where("1=1").Delete(&model.Order{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("清空订单失败: %v", err)
	}

	if err := tx.Model(&model.Profile{}).Unscoped().Where("1=1").Delete(&model.Profile{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("清空用户配置失败: %v", err)
	}

	if err := tx.Model(&model.User{}).Unscoped().Where("1=1").Delete(&model.User{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("清空用户失败: %v", err)
	}

	if err := tx.Model(&model.Product{}).Unscoped().Where("1=1").Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("清空产品失败: %v", err)
	}

	if err := tx.Model(&model.Role{}).Unscoped().Where("1=1").Delete(&model.Role{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("清空角色失败: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %v", err)
	}

	return nil
}
