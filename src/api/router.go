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
		v1.POST("/user/token", func(c *gin.Context) {
			token := c.Query("token")
			fmt.Println(token)
		})		
		v1.DELETE("/user/token",func(c *gin.Context){
			
		})		
		v1.GET("/user/info",func(c *gin.Context){
			
		})		
		v1.GET("/user/token/logout",func(c *gin.Context){
			
		})		
		v1.POST("/user/regisitor",func(c *gin.Context){
			
		})		
		v1.GET("/mp3s",func(c *gin.Context){
			
		})		
		v1.GET("/mp3s/:id/link",func(c *gin.Context){
			
		})		
		v1.POST("/mp3s",func(c *gin.Context){

		})

	}
	config.service.router.Run(":8081")
}