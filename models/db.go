package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func init() {
	var err error
	//dsn := "root:root@tcp(127.0.0.1:3306)/shop_user_srv?charset=utf8&parseTime=True&loc=Local&collation=utf8_general_ci"
	dsn := "gorm:123456@tcp(192.168.0.188:3306)/gorm?charset=utf8&parseTime=True&loc=Local&collation=utf8_general_ci"
	// dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&collation=utf8_general_ci",
	// 	c.User,
	// 	c.Password,
	// 	c.Host,
	// 	c.Port,
	// 	c.DbName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}
}
