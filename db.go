package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func initDB() {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8"
	var err error
	logrus.Info("path:", path)
	for {
		db, err = gorm.Open("mysql", path)
		if err != nil {
			logrus.Error("DB connect error:", err, " Retry in 2 seconds!")
			time.Sleep(time.Second * 2)
			continue
		}
		logrus.Info("DB connect successful!")
		break
	}
	db.LogMode(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "blog_" + defaultTableName
	}
	db.AutoMigrate(&User{})
}
