package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config 配置文件结构体
type Config struct {
	RedisHost     string //redis地址
	RedisPort     string //redis端口
	RedisPassword string //redis密码
	RedisDB       int    //redis数据库 0-15
	Redis         bool   //是否开启redis

	DBHost     string //DB地址
	DBPort     string //DB端口
	DBUsername string //DB用户
	DBPassword string //DB密码
	DBName     string //DB数据库
	DB         bool   //是否开启db
}

// TomlConfig toml的配置文件
type TomlConfig struct {
	Runmode Runmode
	Pro     Pro
	Dev     Dev
}

type Pro = Config
type Dev = Config

type Runmode struct {
	Runmode string
}

var conf = new(Config)

func initConfig() {
	viper.SetConfigName("config")         //  设置配置文件名 (不带后缀)
	viper.AddConfigPath("/etc/appname/")  // 第一个搜索路径
	viper.AddConfigPath("$HOME/.appname") // 可以多次调用添加路径
	viper.AddConfigPath(".")              // 比如添加当前目录
	err := viper.ReadInConfig()           // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	tc := new(TomlConfig)
	if err := viper.Unmarshal(tc); err != nil {
		logrus.Fatal(err)
	}
	if tc.Runmode.Runmode == "dev" {
		*conf = tc.Dev
	} else {
		*conf = tc.Pro
	}
	tc = nil
	logrus.Infof("%+v", conf)
	logrus.Info("初始化config成功")
}
