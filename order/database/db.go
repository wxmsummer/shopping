package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/config"
	log "github.com/micro/go-micro/v2/logger"
	"shopping/order/model"
)

func InitDB() (*gorm.DB, error){
	//加载配置项
	// err := config.LoadFile("./config.json")
	err := config.LoadFile("./order/database/config.json")
	if err != nil {
		log.Fatalf("Could not load config file: %s", err.Error())
		return nil, err
	}
	conf := config.Map()

	//db
	db, err := CreateConnection(conf["mysql"].(map[string]interface{}))

	db.AutoMigrate(&model.Order{})

	if err != nil {
		log.Fatalf("connection error : %v \n" , err)
		return nil, err
	}

	return db, err
}

// db服务：db初始化：如果user表不存在则创建，然后连接数据库
func CreateConnection(conf map[string]interface{}) (*gorm.DB, error) {
	host := conf["host"]
	port := conf["port"]
	user := conf["user"]
	dbName := conf["database"]
	password := conf["password"]
	return gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, dbName,
	),
	)
}

