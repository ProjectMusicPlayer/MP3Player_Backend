package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
//	"net/http"
)

func routerInit(){
	config.service.router = gin.Default()	
	v1 := config.service.router.Group("v1")
	{
		//stu 对接apihdu
		v1.GET("/user/logout", func(c *gin.Context) {
			token := c.Query("token")
			fmt.Println(token)
		})		

	}
	config.service.router.Run(":8081")
}