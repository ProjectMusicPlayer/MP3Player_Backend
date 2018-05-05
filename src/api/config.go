package main

import (
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
)

type Config struct{
	service Service
	appSecret string
	appId string
}

type Service struct {
	router *gin.Engine
	db Sqlconn
	dbBind Sqlconn
	dbAuth Sqlconn
	redis redis.Conn
}

var config Config
var tables Tables
func (c *Config) setDb(){
	//数据库配置
	c.service.db.config.dbname = ""
	c.service.db.config.host = ""
	c.service.db.config.user = ""
	c.service.db.config.pswd = ""
}

func (t *Tables) setTable(){
	t.user = "user"
	t.token = "token"
	t.data = "data"
	t.mp3 = "mp3"
}

type Tables struct{
	user string
	token string
	data string
	mp3 string
}


