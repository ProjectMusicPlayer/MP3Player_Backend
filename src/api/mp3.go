package main

import (
	"strconv"
)

func addMp3(name string,singer string,books string,len string,mp3addr string,lrcaddr string)(int,interface{}){
	len6,err := strconv.Atoi(len)
	if(err!=nil){
		return makeErrJson(403,40300,err)
	}
	leni := int64(len6)
	err = addMp3DB(name,singer,books,leni,lrcaddr,mp3addr)
	if(err!=nil){
		return makeErrJson(403,40301,err)
	}
	return makeSuccessJson("add mp3 success")
}

func searchMp3(key string,token string)(int,interface{}){
	err :=checkLoginToken(token)
	if err!=nil {
		return makeErrJson(401,40101,"unauthorized")
	}
	err = pregCheckChinese(key)
	if err!=nil {
		return makeErrJson(403,40311,err)
	}
	rs,length,err := readMp3Data(key)
	if err!=nil {
		return makeErrJson(403,40312,err)
	}
	data := make(map[string]interface{})
	data["error"] = 0
	data["mag"] = "get mp3 success"
	data["data"] = rs
	data["arrayLength"] = length
	return 0,data
}