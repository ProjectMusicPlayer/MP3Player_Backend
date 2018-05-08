package main

import (
	"fmt"
	"time"
)

func crtStateId()string{
	return  GetRandomString(44)
}

func writeState(state,user,data1,data2,data3 string)(error){
	t := time.Now().Unix()+3600
	_,err:=config.service.db.conn.Exec("insert into state values(?,?,?,?,?,?)",state,user,data1,data2,data3,t)
	if(err!=nil){
		return err
	}
	return nil
}

func readState(state string)(string,string,string,string,int64,error,int){
	rows,err:=config.service.db.conn.Query("select * from state where state = ?",state)
	if(err!=nil){
		return "","","","",0,err,41000
	}
	defer rows.Close()
	if(rows.Next()){
		var p1,p2,p3,p4,p5 string
		var t int64
		err = rows.Scan(&p1,&p2,&p3,&p4,&p5,&t)
		if(err!=nil){
			return "","","","",0,err,41001
		}
		if(time.Now().Unix()<t){
			return p2,p3,p4,p5,t,nil,0
		}
		return "","","","",0,fmt.Errorf("state out of data"),41002
	}else{
		return "","","","",0,fmt.Errorf("invaild stateid"),41003
	}
}