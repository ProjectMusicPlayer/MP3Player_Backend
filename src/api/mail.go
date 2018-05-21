package main

import (
	"fmt"
    "github.com/mikemintang/go-curl"
)
//敏感数据
//var mail_user = *****
//var mail_key = *****

type mailsend struct{
	to string
	title string
	body string
}
func sendmail(to,user,state string)(error){
	req := curl.NewRequest()
	//https://api.mp3.h-00.com/v1/user/regisitor/mailRedirect/
    resp, err := req.
        SetUrl("https://projm.billsjc.com/index.php/console/user_sendmail?addr="+to+"&user="+user+"&&url=https%3a%2f%2fapi.mp3.h-00.com%2fv1%2fuser%2fregisitor%2fmailRedirect%2f"+state).
        Post()

    if err != nil {
        return err
    } else {
        if resp.IsOk() {
			fmt.Print(resp.Body)
			return nil
		}else{
			return fmt.Errorf("send smail failed")
		}
	}
}
func sendmailf(to,user,state string)(error){
	req := curl.NewRequest()
	//https://api.mp3.h-00.com/v1/user/regisitor/mailRedirect/
    resp, err := req.
        SetUrl("https://projm.billsjc.com/index.php/console/user_sendmail2?addr="+to+"&user="+user+"&&url=https%3a%2f%2fapi.mp3.h-00.com%2fv1%2fuser%2fregisitor%2fmailRedirect%2f"+state).
        Post()

    if err != nil {
        return err
    } else {
        if resp.IsOk() {
			fmt.Print(resp.Body)
			return nil
		}else{
			return fmt.Errorf("send smail failed")
		}
	}
}