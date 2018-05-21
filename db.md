# 数据库表
### user 用户表
- `username` **(P)** string,用户名
- `password` string,密码(哈希)
- `regidate` int,注册日期

### token token表
- `username` string,用户名
- `token` **(P)** string,token
- `vaild time` int,有效期

### mp3 歌曲表
- `id` **(P)** int,歌曲id
- `name` string,歌手
- `singer` string,歌曲名
- `books` string,专辑
- `length` int,长度

### mp3_link 链接表
- `id` **(P)** int,歌曲id
- `link` string,下载链接

### data 日志表
- `id` **(AI/P)** int,日志id
- `token` string,使用的token
- `username` string,操作用户名
- `operation` string,操作名
- `time` int 操作时间

### report 建议反馈表
- `id` string,反馈id(工单号)
- `username` string,用户名
- `reportBody` string,反馈主体
- `timeStamp` int,反馈时间

### state 状态表
- `state` **(P)** string,stateid
- `username` string,用户名
- `data1` string,附加参数
- `data2` string,附加参数
- `data3` string,附加参数
- `vaildtime` int,有效期