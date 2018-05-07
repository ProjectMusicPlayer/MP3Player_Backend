package main

import (
	"fmt"
)

func user_regisiter(user string,pswd string)(code int,m map[string]string){
	code = 200
	m = make(map[string]string)
	err := pregCheck2(user,pswd,false)
	if(err!=nil){
		m["error"] = "40001"
		m["msg"] = fmt.Sprint(err)
		return
	}
	err = addUser(user,pswd)
	if(err!=nil){
		m["error"] = "40002"
		m["msg"] = fmt.Sprint(err)
		return
	}
	token,err := writeToken(user)
	if(err!=nil){
		m["error"] = "40003"
		m["msg"] = fmt.Sprint(err)
	}
	m["error"] = "0"
	m["msg"] = "regisitor success"
	m["access_token"] = token
	return
}