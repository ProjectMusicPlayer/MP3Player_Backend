package main

import (
	"fmt"
	"time"
)

func addUser(user string,passwd string,email string)(err error){
	rows,err := config.service.db.conn.Query("select * from  "+tables.user+"  where username = ?",user)
	if(err!=nil){
		return
	}
	defer rows.Close()
	if !(rows.Next()){
		pswdmd5 := md5_encode(passwd)
		timer := time.Now().Unix()
		_,err = config.service.db.conn.Exec("insert into "+tables.user+"  values (?,?,?,?)",user,pswdmd5,email,timer)
		if(err!=nil){
			return
		}else{
			return nil
		}
	}else{
		return fmt.Errorf("user already existe")
	}
}

func checkUser(user string)(b bool,err error){
	rows,err := config.service.db.conn.Query("select * from  "+tables.user+"  where username = ?",user)
	if(err!=nil){
		return false,err
	}
	defer rows.Close()
	if (rows.Next()){
		return true,nil
	}else{
		return false,nil
	}
}


func checkLoginUser(user string,passwd string)(err error){
	rows,err := config.service.db.conn.Query("select * from "+tables.user+" where username = ?",user)
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


func checkLoginToken(token string)(user string,err error){
	rows,err:=config.service.db.conn.Query("select * from token where token = ?",token)
	if(err!=nil){
		return "",err
	}
	defer rows.Close()
	if(rows.Next()){
		var p1,p2 string
		var p3 int64
		err = rows.Scan(&p1,&p2,&p3)
		if(err!=nil){
			return "",err
		}
		if(time.Now().Unix()>p3){
			return "",fmt.Errorf("token timeout")
		}
		return p1,nil
	}
	return "",fmt.Errorf("invaild token")
}

func writeToken(user string)(token string,err error){
	token = tokenCrt()
	timev := int(time.Now().Unix())
	_,err=config.service.db.conn.Exec("insert into "+tables.token+" values(?,?,?)",user,token,timev)
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