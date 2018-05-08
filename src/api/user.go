package main

import (
	"fmt"
)

func user_regisiter(user string,pswd string,email string)(code int,m map[string]interface{}){
	code = 200
	m = make(map[string]interface{})
	err := pregCheck2(user,pswd,false)
	if(err!=nil){
		m["error"] = 40001
		m["msg"] = fmt.Sprint(err)
		return
	}
	b,err := checkUser(user)
	if(err!=nil){
		m["error"] = 40002
		m["msg"] = fmt.Sprint(err)
		return
	}
	if(b){
		m["error"] = 40003
		m["msg"] = "user already exist"
		return
	}
	state := crtStateId()
	err = writeState(state,user,pswd,email,"")
	if(err!=nil){
		m["error"] = 40004
		m["msg"] = fmt.Sprint(err)
	}
	err = sendmail(email,user,state)
	m["error"] = 0
	m["msg"] = "mail send success"
	return
}

func user_regisiter_mailRedirect(state string)(code int,data string){
	err:=pregCheck(state,false)
	if(err!=nil){
		code = 200
		data =  fmt.Sprint(err)
		return
	}
	username,passwd,mailaadr,_,_,err,datacode:=readState(state)
	if(err!=nil){
		code = 200
		data =  fmt.Sprint(err)
		return
	}
	if(datacode!=0){
		code = 200
		data = fmt.Sprintf("return err code %d",datacode)
		return
	}
	err = addUser(username,passwd,mailaadr)
	if(err!=nil){
		code = 200
		data =  fmt.Sprint(err)
		return
	}
	code = 200
	data = "regisiter success"
	return
}

func user_login(user,passwd string)(int,map[string]interface{}){
	err := pregCheck2(user,passwd,false)
	if(err!=nil){
		return makeErrJson(42000,err)
	}
	err = checkLoginUser(user,passwd)
	if(err!=nil){
		return makeErrJson(42001,err)
	}
	token,err := singToken(user)
	if(err!=nil){
		return makeErrJson(42002,err)
	}
    var m map[string]interface{}
    m = make(map[string]interface{})
    m["error"] = 0
	m["msg"] = "login success"
	m["token"] = token
    return 200,m
}