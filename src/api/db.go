package main

/*
@Author SJC
@LastEdit 18.4.19
简单的数据库封装，主要实现了连接，增删改优化，1-4返回值查询封装
*/

import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

/*
class Sqlconn
其实这个东西没有什么用处的
*/

type Sqlconn struct{
	//stus 0-未初始化 1-设置成功(未连接) 2-连接成功(已连接) 3-断开连接(配置可连接) -1-(连接失败) -2-出现错误
	config *DbConfig
	stus int
	conn *sql.DB	//数据库结构指针
}

func (s *Sqlconn) set (conf *DbConfig){
	s.config = conf
	s.stus = 1
}

func (s *Sqlconn) start () bool{
	if(s.stus != 0){
		str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",s.config.user,s.config.pswd,s.config.host,s.config.dbname)
		db , err := sql.Open("mysql",str)
		if err != nil {
			s.stus = -1
			return false
		}else{
			s.stus = 2
			s.conn = db
			return true
		}
	}else{
		fmt.Println("Sqlconn is not set!")
	}
	return false
}

func (s *Sqlconn) close (){
	s.conn.Close()
	s.stus = 3
}


/*
class DbConfig
数据库配置文件
*/
type DbConfig struct{
	host string	//主机地址+端口(或域名)
	user string	//数据库用户名
	pswd string	//数据库密码
	dbname string	//数据库名
}

//数据库启动封装，根据输入选择数据库
func initDB(ser string)bool{
	//测试服务器
	testDB := &DbConfig{
		"localhost:3306",
		"mp3test",
		"mp3test",
		"mp3test",
	}
	//测试服务器
	testDBR := &DbConfig{
		"112.124.47.125:3306",
		"mp3test",
		"mp3test",
		"mp3test",
	}
	//运营服务器
	releaseDB := &DbConfig{
		"mp3test",
		"mp3test",
		"mp3test",
		"mp3test",
	}

	//选择数据库
	switch (ser){
	case "develop" : config.service.db.set(testDB)
	case "developR" : config.service.db.set(testDBR)
	case "release" : config.service.db.set(releaseDB)
	}

	b1 := config.service.db.start()
	return b1
}