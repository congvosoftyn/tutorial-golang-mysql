package database

import (
	"log"

	"github.com/vothecong/go-tutorial/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectionDB() {

	dsn := "root:123456@tcp(localhost:3306)/DBDemo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatal("Connect mysql faild!")
	}

	log.Println("Connect MySQL success!")
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.Book{})

	DB = db
}
