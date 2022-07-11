package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"app-produk/entity"
)

func InitDB() *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func MigrateDB(conn *gorm.DB)  {
	conn.AutoMigrate(entity.Product{})
}