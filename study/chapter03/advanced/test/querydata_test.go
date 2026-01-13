package test

import (
	"local/go_study/study/chapter03/advanced/db"
	"local/go_study/study/chapter03/advanced/model"
	"testing"
)

func TestQueryDataProducts(t *testing.T) {

	DB, err := db.DBInit()
	if err != nil {
		t.Fatalf("数据库连接失败: %v", err) // 连接失败直接终止测试，打印具体错误
	}

	var products []model.Product

	err = DB.Find(&products).Error
	if err != nil {
		t.Errorf("查询产品失败: %v", err) // 查询失败标记为测试错误，不终止
		return
	}

	if len(products) == 0 {
		t.Log("⚠️ 未查询到任何Product数据")
		return
	}

	t.Logf("✅ 成功查询到 %d 个Product数据：", len(products))
	for idx, product := range products {
		t.Logf("  产品%d, 名称=%s, 价格=%d元, SKU=%s",
			idx+1, product.Name, product.Price, product.SKU)
	}
}

func TestQueryDataUser(t *testing.T) {
	DB, err := db.DBInit()
	if err != nil {
		t.Fatalf("数据库连接失败: %v", err) // 连接失败直接终止测试，打印具体错误
	}

	var users []model.User

	err = DB.
		Preload("Profile").
		Preload("Roles").
		Preload("Orders").
		Preload("Orders.Items.Product").
		Find(&users).
		Error
	if err != nil {
		t.Errorf("查询数据失败: %v", err)
		return
	}

	if len(users) == 0 {
		t.Logf("没有查询到数据")
		return
	}

	t.Logf("✅ 成功查询到 %d 个用户，关联信息如下：", len(users))
	for userIdx, user := range users {
		t.Logf("\n===== 用户 %d 核心信息 =====", userIdx+1)
		t.Logf("ID: %d | 姓名: %s | 邮箱: %s", user.ID, user.Name, user.Email)

		// 打印关联的Profile（一对一）
		if user.Profile.ID > 0 {
			t.Logf("【档案】昵称: %s | 手机号: %s | 地址: %s",
				user.Profile.Nickname, user.Profile.Phone, user.Profile.Address)
		} else {
			t.Logf("【档案】无关联档案")
		}

		// 打印关联的Roles（多对多）
		if len(user.Roles) > 0 {
			roleNames := make([]string, 0, len(user.Roles))
			for _, role := range user.Roles {
				roleNames = append(roleNames, role.Name)
			}
			t.Logf("【角色】%v", roleNames)
		} else {
			t.Log("【角色】无关联角色")
		}

		// 打印关联的Orders + 嵌套的OrderItems + Product（一对多+嵌套）
		if len(user.Orders) > 0 {
			t.Logf("【订单】共 %d 个订单：", len(user.Orders))
			for orderIdx, order := range user.Orders {
				t.Logf("订单%d：编号=%s | 状态=%s | 总价=%d元",
					orderIdx+1, order.OrderNo, order.Status, order.TotalPrice)

				// 打印订单项及关联商品
				if len(order.Items) > 0 {
					t.Log("订单项：")
					for itemIdx, item := range order.Items {
						t.Logf("项%d：商品=%s | 单价=%d元 | 数量=%d | SKU=%s",
							itemIdx+1, item.Product.Name, item.Product.Price, item.Quantity, item.Product.SKU)
					}
				} else {
					t.Log("    订单项：无")
				}
			}
		} else {
			t.Log("【订单】无关联订单")
		}
	}

}
