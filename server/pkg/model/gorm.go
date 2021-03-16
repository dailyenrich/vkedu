package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"vke/pkg/config"
)

func InitGormDB() *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				config.Get().Mysql.Username,
				config.Get().Mysql.Password,
				config.Get().Mysql.Hostname,
				config.Get().Mysql.Port,
				config.Get().Mysql.Database,
			),
		),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	//设置连接池
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
