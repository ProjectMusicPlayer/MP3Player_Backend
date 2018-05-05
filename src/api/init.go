package main

import (
	"fmt"
//	"github.com/gin-gonic/gin"
)

func initService(){
	//*
	//初始化sql
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