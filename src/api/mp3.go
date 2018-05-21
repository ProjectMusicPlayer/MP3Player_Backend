package main

import (
	"strconv"
)

func addMp3(name string,singer string,books string,len string,mp3addr string,lrcaddr string)(int,interface{}){
	len6,err := strconv.Atoi(len)
	leni := int64(len6)
	if(err!=nil){
		return makeErrJson(403,40300,err)
	}
	err = addMp3DB(name,singer,books,leni,mp3addr,lrcaddr)
	if(err!=nil){
		return makeErrJson(403,40301,err)
	}
	return makeSuccessJson("add mp3 success")
}