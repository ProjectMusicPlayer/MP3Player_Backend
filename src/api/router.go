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
		v1.GET("/user/login", func(c *gin.Context) {
			token := c.Query("token")
			fmt.Println(token)
		})		
		v1.GET("/user/logout",func(c *gin.Context){
			
		})		
		v1.GET("/user/info",func(c *gin.Context){
			
		})		
		v1.GET("/user/token/logout",func(c *gin.Context){
			
		})		
		v1.GET("/user/regisitor",func(c *gin.Context){
			
		})		
		v1.GET("/mp3/serch",func(c *gin.Context){
			
		})		
		v1.GET("/user/downloadLink",func(c *gin.Context){
			
		})		
		v1.POST("/user/downloadLink",func(c *gin.Context){
			
		})

	}
	config.service.router.Run(":8081")
}