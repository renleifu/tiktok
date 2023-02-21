package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func init() {

	driverName := AppConfig.Get("datasource.driverName").(string)
	dataSourceName := AppConfig.Get("datasource.dataSourceName").(string)

	//打印文件读取出来的内容：
	log.Printf("数据库为 %s,数据库链接为%s", driverName, dataSourceName)
	db, err := gorm.Open(driverName, dataSourceName+"?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	DB = db
}
