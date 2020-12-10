package common

import (
	"ginessential/model"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDB() *gorm.DB  {
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "hxy116991"
	charset := "utf8"

	//dsn := "root:hxy116991@tcp(127.0.0.1:3306)/ginessential?charset=utf8&parseTime=True&loc=Local"
	dsn := username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset="+charset+"&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB  {
	return DB
}