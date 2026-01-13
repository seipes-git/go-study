package basic

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Global_DB *gorm.DB

func main() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:               "Alan:seipes@tcp(127.0.0.1:3306)/gostudy?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize: 191,                                                                             // default size for string fields                                                                         // auto configure based on currently MySQL version
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gva_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	Global_DB = db

	TestUserCreate()

	// CreatedSingleTest()
	// CreatedMultipleTest()

	// FindTest()

	// UpdateTest()

	// DeleteTest()

	// type User struct {
	// 	Name string
	// }

	// M := db.Migrator()
	// M.CreateTable(&User{})
	// fmt.Println(M.HasTable(&User{}))
	// fmt.Println(M.HasTable("gva_users"))
	// fmt.Println(M.DropTable(&User{}))
	// M.RenameTable(&User{}, "gva_user_two")
}
