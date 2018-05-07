package main

import (
	"github.com/gin-gonic/gin"
//	"net/http"
)

func routerInit(){
	config.service.router = gin.Default()	
	v1 := config.service.router.Group("v1")
	{
		v1.POST("/user/token", func(c *gin.Context) {
			
		})		
		v1.DELETE("/user/token",func(c *gin.Context){
			
		})		
		v1.GET("/user/info",func(c *gin.Context){
			
		})		
		v1.GET("/user/token/logout",func(c *gin.Context){
			
		})		
		v1.POST("/user/regisitor",func(c *gin.Context){
			c.JSON(user_regisiter(c.Query("username"),c.Query("password")))
		})		
		v1.GET("/mp3s",func(c *gin.Context){
			
		})		
		v1.GET("/mp3s/:id/link",func(c *gin.Context){
			
		})		
		v1.POST("/mp3s",func(c *gin.Context){

		})

	}
	//mp3服务挂靠在8082端口
	config.service.router.Run(":8082")
}