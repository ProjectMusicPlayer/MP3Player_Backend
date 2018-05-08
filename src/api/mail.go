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
        SetUrl("************").
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



/*
参数	类型	必须	说明
apiUser	string	是	API_USER
apiKey	string	是	API_KEY
from	string	是	发件人地址. 举例: support@ifaxin.com, 为了更高的送达率，建议from域名后缀与发信域名一致。
to	string	*	收件人地址. 多个地址使用';'分隔, 如 ben@ifaxin.com;joe@ifaxin.com
subject	string	是	标题. 不能为空
html	string	*	邮件的内容. 邮件格式为 text/html
contentSummary	string	*	邮件摘要. 该字段传入值后，若原邮件已有摘要，会覆盖原邮件的摘要；若原邮件中没有摘要将会插入摘要。了解邮件摘要的更多内容，请点击这里
fromName	string	否	发件人名称. 显示如: ifaxin客服支持<support@ifaxin.com>
cc	string	否	抄送地址. 多个地址使用';'分隔
bcc	string	否	密送地址. 多个地址使用';'分隔
replyTo	string	否	设置用户默认的回复邮件地址.多个地址使用';'分隔，地址个数不能超过3个. 如果 replyTo 没有或者为空, 则默认的回复邮件地址为 from
labelId	int	否	本次发送所使用的标签ID. 此标签需要事先创建
headers	string	否	邮件头部信息. JSON 格式, 比如:{"header1": "value1", "header2": "value2"}
attachments	file	否	邮件附件. 发送附件时, 必须使用 multipart/form-data 进行 post 提交 (表单提交)
xsmtpapi	string	否	SMTP 扩展字段. 详见 X-SMTPAPI
plain	string	否	邮件的内容. 邮件格式为 text/plain
respEmailId	string (true, false)	否	默认值: true. 是否返回 emailId. 有多个收件人时, 会返回 emailId 的列表
useNotification	string (true, false)	否	默认值: false. 是否使用回执
useAddressList	string (true, false)	否	默认值: false. 是否使用地址列表发送. 比如: to=group1@maillist.sendcloud.o
*/