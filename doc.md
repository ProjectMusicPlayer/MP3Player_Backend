# MP3Player 接口文档


- 适用范围 : `unknow`
- baseurl : `https://api.mp3.h-00.com`
- version:v1

## 访问方式:
```
{baseurl}/{version}/{method}?{params}
e.g:
https://api.mp3.h-00.com/v1/debug/user/login?username=......&password=......
```

## 接口认证
API 采用了 OAuth2 验证机制 (See [OAuth](https://en.wikipedia.org/wiki/OAuth), [RFC6749](https://tools.ietf.org/html/rfc6749)), 我们提供了 `AuthorizationCode Flow` 和 `ClientCredentials Flow`, 根据应用需求发放 Flow 权限.<br>
在获取到 `AccessToken` 之后, 使用请求头 `Authorization` (Previously `X-Access-Token`, `Authorization` is preferred.) 携带 Token ,并使用请求头`secret`携带应用密钥,访问 API.

## FAQ
- 使用此接口，即代表调用者默认同意 [The JSON License](http://www.json.org/license.html)
- 接口入口使用正则匹配，请确保完全填满所需数据，否则会被`401 invaild params`
## POST `/user/token` 用户登录
发送用户名密码完成登录，若成功则返回Access_token<br>
Params:
- `username` string,用户名
- `password` string,密码
- `timestamp` string,发送请求的unix时间

返回结果:<br>
成功:
```
{
    error:"0",   //成功代码0
    msg:"login success", //成功消息
    access_token:"{access_token}",   //用户的token
    invaildtime:"{timestamp}"   //token有效期，一般为7天
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```
## DELETE `/user/token` 用户登出
发送token进行注销，注销之后token立即失效<br>

成功:
```
{
    error:"0",   //成功代码0
    msg:"logout success", //成功消息
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```
## GET `/user/info` 获取用户信息


成功:
```
{
    error:"0",   //成功代码0
    msg:"get user info success", //成功消息
    usrname:"{username}"    //用户名
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```

## GET `/user/token/info` 获取token过期信息

发送token来获取当前token有效期<br>
**注意**:过期token将会在过期之后很短时间内消失并无法查询



成功:
```
{
    error:"0",   //成功代码0
    msg:"get token status success", //成功消息
    queryTimeStamp:"{timestamp}", //查询时间
    vaild:"{vaild}",    //token状态 1有效，2无效
    vaildTime:"{timestamp}" //token过期时间
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```
## POST `/user/regisitor` 用户注册
用户通过此接口注册，将会发送一封激活邮件<br>

Params:
- `username` string,用户名 
- `email` string,邮箱地址

成功:
```
{
    error:"0",   //成功代码0
    msg:"send mail success", //成功消息
    vaildTime:"{timestamp}" //mail过期时间
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```
## GET `/user/regisitor/mail` 邮箱注册链接接口
## POST `/user/forget` 忘记密码
用户通过此接口注册，将会发送一封重置密码邮件<br>

Params:
- `username` string,用户名
- `邮箱` string,用户名

成功:
```
{
    error:"0",   //成功代码0
    msg:"send mail success", //成功消息
    vaildTime:"{timestamp}" //mail过期时间
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```
## POST `/user/password` 修改密码
用户通过此接口注册，将会发送一封重置密码邮件<br>

Params:
- `old` string,用户名
- `new` string,用户名

成功:
```
{
    error:"0",   //成功代码0
    msg:"send mail success", //成功消息
    vaildTime:"{timestamp}" //mail过期时间
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```
## GET `/user/forget/mail` 邮箱找回密码链接接口
## GET `/mp3s` 获取搜索结果列表
发送搜索关键词获取搜索结果数组<br>

Params:
- `keyword` string,歌曲关键词

成功:
```
{
    error:"0",   //成功代码0
    msg:"get token status success", //成功消息
    data:[  //返回结果数组(无数据则返回空数组)
        {
            id:"{id}"   //歌曲id
            name:"{name}"   //歌名
            singer:"{singer}"   //歌手
            books:"{books}" //专辑
            length:"{length}"   //歌曲长度
        },
        {
            id:"{id}"   //歌曲id
            name:"{name}"   //歌名
            singer:"{singer}"   //歌手
            books:"{books}" //专辑
            length:"{length}"   //歌曲长度
        },
        {
            id:"{id}"   //歌曲id
            name:"{name}"   //歌名
            singer:"{singer}"   //歌手
            books:"{books}" //专辑
            length:"{length}"   //歌曲长度
        }
    ]
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```
## GET `/mp3s/{mp3id}/link` 获取下载链接

通过搜索结果的歌曲id返回cdn下载链接<br>

Params:
- `mp3id` string,歌曲id

成功:
```
{
    error:"0",   //成功代码0
    msg:"get token status success", //成功消息
    url:"{url}",    //下载列表
    size:"{size}"   //文件大小
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```
## POST `/mp3s` 提交新的歌曲
**此接口Access_token不可用**  后台提交歌曲信息

<br>

Params:
- `name` 歌名
- `singer` 歌手
- `books` 专辑
- `length` 歌曲长度
- `url` 下载地址

成功:
```
{
    error:"0",   //成功代码0
    msg:"update song info success", //成功消息
    id:"{mp3id}"    //上传的mp3 id
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```

## PUT `/report` 提交工单
用户在反馈页面提交工单
<br>

Params:
- `token` 用户token
- `body` 反馈主体

成功:
```
{
    error:"0",   //成功代码0
    msg:"report success", //成功消息
}
```
错误:
```
{
    error:"{error_code}",   //错误代码
    msg:"{error_massage}" //错误消息
}
```


---
- `@Autor` : `SJC`
- `@version` : `1.1`
- `@lastUpdate` : `18.5.5 14:03`
- [The JSON License](http://www.json.org/license.html)