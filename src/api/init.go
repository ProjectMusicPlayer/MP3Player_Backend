package main

import (
	"fmt"
//	"github.com/gin-gonic/gin"
)

func initService(){
	//*
	//初始化sql
	//初始化数据表
	tables.setTable()
	if(initDB("develop")){
		fmt.Println("load config successfully")
		//初始化gin
		routerInit()
	}else{
		fmt.Println("load config failed")
	}
	//*/
	//routerInit()
}