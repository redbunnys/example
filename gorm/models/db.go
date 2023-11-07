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

type Article struct {
	// 文章ID
	ArticleId int `gorm:"primary_key" json:"articleId"`
	// 文章标题
	ArticleTitle string `json:"articleTitle"`
	// 文章描述
	Tags       []Tag `gorm:"many2many:article_tag;foreignKey:ArticleId;joinForeignKey:ArticleId;References:TagId;joinReferences:TagId" json:"tags"`
	CategoryId int   `json:"categoryId"`
}

type Tag struct {
	// 标签ID
	TagId int `gorm:"primary_key" json:"tagId"`
	// 标签名称
	TagName  string    `json:"tagName"`
	Articles []Article `gorm:"many2many:article_tag;foreignKey:TagId;joinForeignKey:TagId;References:ArticleId;joinReferences:ArticleId" json:"articles"`
}
type Category struct {
	// 分类ID
	CategoryId int `gorm:"primary_key" json:"categoryId"`
	// 分类名称
	CategoryName string    `json:"categoryName"`
	Articles     []Article `gorm:"foreignKey:CategoryId""`
}
