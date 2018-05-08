package main

import (
	"time"
)

func tokenCrt()string{
	return GetRandomString(44)
}

func singToken(user string)(token string,err error){
	token = tokenCrt()
	timen := time.Now().Unix()
	_,err = config.service.db.conn.Exec("insert into token values(?,?,?)",user,token,timen)
	if(err!=nil){
		return "",err
	}
	return token,nil
}