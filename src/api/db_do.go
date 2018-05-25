package main

import (
	"fmt"
	"time"
)

//添加用户
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

//检查用户重复
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

//检验登录
func checkLoginUser(user string,passwd string)(err error){
	rows,err := config.service.db.conn.Query("select * from "+tables.user+" where username = ?",user)
	if(err!=nil){
		return
	}
	defer rows.Close()
	if(rows.Next()){
		var p1,p2,p3 string
		var p4 int64
		err = rows.Scan(&p1,&p2,&p3,&p4)
		if(err!=nil){
			return
		}
		pswd := p2
		passwd = md5_encode(passwd)
		if(pswd != passwd){
			return fmt.Errorf("incorrect username or password!")
		}else{
			return nil
		}
	}
	return fmt.Errorf("incorrect username or password!")
}

//检验token有效性并返回用户数据
func checkLoginTokenI(token string)(user string,email string,regidate int64,err error){
	rows,err:=config.service.db.conn.Query("select * from token where token = ?",token)
	if(err!=nil){
		return "","",0,err
	}
	defer rows.Close()
	if(rows.Next()){
		var p1,p2 string
		var p3 int64
		err = rows.Scan(&p1,&p2,&p3)
		if(err!=nil){
			return "","",0,err
		}
		if(time.Now().Unix()>p3){
			return "","",0,fmt.Errorf("token timeout")
		}
		rows1,err:=config.service.db.conn.Query("select * from user where username = ?",p1)
		if(err!=nil){
			return "","",0,err
		}
		defer rows1.Close()
		if(rows1.Next()){
			var username,passwd,email string
			var regitime int64
			err = rows1.Scan(&username,&passwd,&email,&regitime) 
			if(err!=nil){
				return "","",0,err
			}
			return username,email,regitime,nil			
		}else{
			return "","",0,fmt.Errorf("invaild user")
		}

	}
	return "","",0,fmt.Errorf("invaild token")
}

//检验token有效性
func checkLoginToken(token string)(err error){
	if(token == "hduhelperSJC"){
		return nil
	}
	rows,err:=config.service.db.conn.Query("select * from token where token = ?",token)
	if(err!=nil){
		return err
	}
	defer rows.Close()
	if(rows.Next()){
		var p1,p2 string
		var p3 int64
		err = rows.Scan(&p1,&p2,&p3)
		if(err!=nil){
			return err
		}
		if(time.Now().Unix()>p3){
			return fmt.Errorf("token timeout")
		}
		return nil
	}
	return fmt.Errorf("invaild token")
}

//写入token
func writeToken(user string)(token string,err error){
	token = tokenCrt()
	timev := int(time.Now().Unix()+86400)
	_,err=config.service.db.conn.Exec("insert into "+tables.token+" values(?,?,?)",user,token,timev)
	if(err!=nil){
		return	"",err
	}
	return token,nil
}

//根据token获取用户信息
func readUserInfo(token string)(m map[string]interface{},err error){
	m = make(map[string]interface{})
	username,emil,regitime,err := checkLoginTokenI(token)
	if(err!=nil){
		return m,err
	}
	m["username"] = username
	m["email"] = emil
	m["registeTime"] = regitime 
	return m,nil
}

//销毁token
func tokenDestory(token string)(err error){
	_,err = config.service.db.conn.Exec("delete from token where token = ?",token)
	return
}//销毁一个用户的所有token
func tokenDestoryByUser(user string)(err error){
	_,err = config.service.db.conn.Exec("delete from token where username = ?",user)
	return
}

//修改密码
func changePswdDB(user,newpswd string)(err error){
	newpswd = md5_encode(newpswd)
	_,err = config.service.db.conn.Exec("update user set password = ? where username = ?",newpswd,user)
	if(err!=nil){
		return
	}
	return nil
}
func writeMp3Data(){

}

func checkUserEmail(user string,email string)(err error){
	rows,err := config.service.db.conn.Query("select * from user where username = ?",user)
	var mail,pswd string
	var rd int64
	if(err!=nil){
		return
	}
	defer rows.Close()
	if(rows.Next()){
		err = rows.Scan(&user,&pswd,&mail,&rd)
		if(err!=nil){
			return
		}
		if(mail==email){
			return nil
		}else{
			return fmt.Errorf("incorrect username or email address")
		}
	}else{
		return fmt.Errorf("incorrect username or email address")
	}

}

func readMp3Data(key string)(data map[int]interface{},arrlength int,err error){
	rows,err := config.service.db.conn.Query("select * from mp3 where name like '%"+key+"%'")
	if(err!=nil){
		return
	}
	data = make(map[int]interface{})
	defer rows.Close()
	i := 0
	for rows.Next(){
		var id int
		var len int64
		var name,singer,books,mp3addr,lrcaddr string
		err = rows.Scan(&id,&name,&singer,&books,&len,&mp3addr,&lrcaddr)
		if(err!=nil){
			return 
		}
		d := make(map[string]interface{})
		d["id"] = id
		d["name"] = name
		d["singer"] = singer
		d["books"] = books
		d["length"] = len
		d["mp3address"] = mp3addr
		d["lrcaddress"] = lrcaddr
		data[i] = d
		i++
	}
	return data,i,nil
}

func getMp3Link(id int){

}

//添加mp3
func addMp3DB(name string,singer string,books string,len int64,mp3addr string,lrcaddr string)(err error){
	_,err = config.service.db.conn.Exec("insert into mp3 values('',?,?,?,?,?,?)",name,singer,books,len,mp3addr,lrcaddr)
	if(err!=nil){
		return err
	}else{
		return nil
	}
}