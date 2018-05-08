package main

import (
	"github.com/gin-gonic/gin"
//	"net/http"
)

func routerInit(){
	config.service.router = gin.Default()	
	v1 := config.service.router.Group("v1")
	{
		//release

		//login
		v1.POST("/user/token", func(c *gin.Context) {
			c.JSON(user_login(c.Query("username"),c.Query("password")))
		})		
		v1.DELETE("/user/token",func(c *gin.Context){
			
		})		
		v1.GET("/user/info",func(c *gin.Context){
			
		})		
		v1.GET("/user/token/logout",func(c *gin.Context){
			
		})		
		v1.POST("/user/regisitor",func(c *gin.Context){
			c.JSON(user_regisiter(c.Query("username"),c.Query("password"),c.Query("email")))
		})			
		v1.GET("/user/regisitor/mailRedirect/:state",func(c *gin.Context){
			stus,data := user_regisiter_mailRedirect(c.Param("state"))
			if(stus == 200){
				c.String(200,data)
			}else{
				c.Redirect(302,data)
			}
		})		
		v1.GET("/mp3s",func(c *gin.Context){
			
		})		
		v1.GET("/mp3s/:id/link",func(c *gin.Context){
			
		})		
		v1.POST("/mp3s",func(c *gin.Context){

		})

		//debug	
		//regisitor
		v1.GET("/debug/user/regisitor",func(c *gin.Context){
			c.JSON(user_regisiter(c.Query("username"),c.Query("password"),c.Query("email")))
		})
		//login	
		v1.GET("/debug/user/login",func(c *gin.Context){
			c.JSON(user_login(c.Query("username"),c.Query("password")))
		})

	}
	//mp3服务挂靠在8082端口
	config.service.router.Run(":8082")
}