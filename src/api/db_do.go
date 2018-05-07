package main

import (
	"fmt"
	"time"
)

func addUser(user string,passwd string)(err error){
	rows,err := config.service.db.conn.Query("select * from ? where username = ?",tables.user)
	if(err!=nil){
		return
	}
	defer rows.Close()
	if(rows.Next()){
		pswdmd5 := md5_encode(passwd)
		timer := time.Now().Unix()
		_,err = config.service.db.conn.Exec("insert into ? values (?,?,?)",tables.user,user,pswdmd5,timer)
		if(err!=nil){
			return
		}else{
			return nil
		}
	}else{
		return fmt.Errorf("user already existe")
	}
}

func checkLoginUser(user string,passwd string)(err error){
	rows,err := config.service.db.conn.Query("select * from ? where username = ?",tables.user)
	if(err!=nil){
		return
	}
	defer rows.Close()
	if(rows.Next()){
		var p1,p2,p3 string
		err = rows.Scan(&p1,&p2,&p3)
		if(err!=nil){
			return
		}
		pswd := p2
		passwd = md5_encode(pswd)
		if(pswd != passwd){
			return fmt.Errorf("incorrect username or password!")
		}else{
			return nil
		}
	}
	return fmt.Errorf("incorrect username or password!")
}

func checkLoginToken(token string){

}

func writeToken(user string)(token string,err error){
	token = tokenCrt()
	timev := int(time.Now().Unix())
	_,err=config.service.db.conn.Exec("insert into ? values(?,?,?)",tables.token,user,token,timev)
	if(err!=nil){
		return	"",err
	}
	return token,nil
}

func writeMp3Data(){

}

func readMp3Data(id int){

}

func searchMp3(key string){

}

func getMp3Link(id int){

}