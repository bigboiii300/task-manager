package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	models2 "task-manager/models"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	port := viper.Get("DB_PORT")
	host := viper.Get("DB_HOST")
	user := viper.Get("DB_USER")
	password := viper.Get("DB_PASS")
	name := viper.Get("DB_NAME")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, name)
	fmt.Println(dsn)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	_ = db.AutoMigrate(&models2.Task{})
	_ = db.AutoMigrate(&models2.User{})
	return db
}
