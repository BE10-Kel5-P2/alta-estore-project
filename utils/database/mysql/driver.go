package mysql

import (
	"altaproject2/config"
	cartdata "altaproject2/features/cart/data"
	orderdata "altaproject2/features/order/data"
	productdata "altaproject2/features/product/data"
	userdata "altaproject2/features/user/data"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Appconfig) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username,
		cfg.Password, cfg.Address, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		log.Println("Cannot connect to db")
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(userdata.User{}, productdata.Product{}, cartdata.Cart{}, orderdata.Order{})
}
