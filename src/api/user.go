package main

import (
	"fmt"
)


/*
user控制器
@Author SJC
@Update at 18.5.9
*/


func user_regisiter(user string,pswd string,email string)(code int,m map[string]interface{}){
	code = 200
	m = make(map[string]interface{})
	err := pregCheck2(user,pswd,false)
	if(err!=nil){
		return makeErrJson(401,40001,err)
	}
	b,err := checkUser(user)
	if(err!=nil){
		return makeErrJson(401,40002,err)
	}
	if(b){
		return makeErrJson(401,40003,"user already existes")
	}
	state := crtStateId()
	err = writeState(state,user,pswd,email,"")
	if(err!=nil){
		return makeErrJson(401,40004,err)
	}
	err = sendmail(email,user,state)
	return makeSuccessJson("send mail success")
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
//登录并签发token
func user_login(user,passwd string)(int,map[string]interface{}){
	err := pregCheck2(user,passwd,false)
	if(err!=nil){
		return makeErrJson(401,42000,err)
	}
	err = checkLoginUser(user,passwd)
	if(err!=nil){
		return makeErrJson401(42001,err)
	}
	token,err := singToken(user)
	if(err!=nil){
		return makeErrJson(401,42002,err)
	}
    var m map[string]interface{}
    m = make(map[string]interface{})
    m["error"] = 0
	m["msg"] = "login success"
	m["token"] = token
    return 200,m
}

//获取用户信息
func getUserInfo(token string)(int,map[string]interface{}){
	err := pregCheck(token,false)
	if(err!=nil){
		return makeErrJson(401,42100,err)
	}
	data,err := readUserInfo(token)
	if(err!=nil){
		return makeErrJson401(42101,err)
	}
	var m map[string]interface{}
	m = make(map[string]interface{})
    m["error"] = 0
	m["msg"] = "get user info success"
	m["data"] = data
	return 200,m
}

//登出并销毁token
func user_logout(token string)(int,map[string]interface{}){
	err := pregCheck(token,false)
	if(err!=nil){
		return makeErrJson(401,43000,err)
	}
	err = tokenDestory(token)
	if(err!=nil){
		return makeErrJson(401,43000,err)
	}
	return makeSuccessJson("logout success")

}

//修改密码 
func changePswd(token,old,new string)(int,map[string]interface{}){
	err := pregCheck3(token,new,old,false)
	if(err!=nil){
		return makeErrJson401(42204,err)
	}
	user,_,_,err := checkLoginTokenI(token)
	if(err!=nil){
		return makeErrJson401(42200,err)
	}
	err = checkLoginUser(user,old)
	if(err!=nil){
		return makeErrJson401(42201,err)
	}
	if(old==new){
		return makeErrJson401(42202,"old password is same as new password")
	}
	err = changePswdDB(user,new)
	if(err!=nil){
		return makeErrJson401(42202,err)
	}
	err = tokenDestoryByUser(user)
	if(err!=nil){
		return makeErrJson401(42203,err)
	}
	return makeSuccessJson("update password successfully")
}

func forgetPswd(user,email string)(int,interface{}){
	err := pregCheck(user,false)
	if(err!=nil){
		return makeErrJson(401,40110,err)
	}
	state := crtStateId()
	err = writeState(state,user,"forget","","")
	if(err!=nil){
		return makeErrJson(401,40111,err)
	}
	err = sendMailApi2(email,user,state)
	if(err!=nil){
		return makeErrJson(401,40112,err)
	}
	return makeSuccessJson("send mail success")
}

func forgetPswdCallback(new,old,state string)(int,interface{}){
	err := pregCheck2(old,new,false)
	if(err!=nil){
		return makeErrJson(403,40301,err)
	}
	var user,data1 string
	user,data1,_,_,_,err,errcode := readState(state)
	if(data1!="forget"){
		return makeErrJson(403,errcode,"invaild state data")
	}
	if(new!=old||new==""){
		return makeErrJson(401,40103,fmt.Errorf("inviald password"))
	}
	err = changePswdDB(user,new)
	if(err!=nil){
		return makeErrJson(403,40304,err)
	}
	return makeSuccessJson("change password success")
}