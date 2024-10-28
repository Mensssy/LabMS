package main

import (
	"com.mensssy.LabMS/conf"
	"com.mensssy.LabMS/controller/router"
	"com.mensssy.LabMS/dao/db"
)

func main() {
	//初始化
	//配置
	conf.InitWebConf()
	conf.InitDBConf()

	//连接数据库
	db.DBConnect()

	//启动路由
	r := router.GetRouter()
	r.Run(conf.ServerPort)
}
